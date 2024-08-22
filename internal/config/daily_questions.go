package config

import (
	"fmt"
	"sync"
	"time"

	"github.com/rarimo/geo-points-svc/internal/data"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
)

type DailyQuestions struct {
	Timezone    int
	Deadlines   map[string]int64
	Responders  []string
	muDeadlines sync.RWMutex
	muResponses sync.RWMutex
}

func (c *config) DailyQuestions() *DailyQuestions {
	return c.DailyQuestion.Do(func() interface{} {
		var cfg struct {
			Timezone int `fig:"timezone"`
		}

		err := figure.Out(&cfg).
			From(kv.MustGetStringMap(c.getter, "daily_questions")).
			Please()
		if err != nil {
			panic(fmt.Errorf("failed to figure out daily questions config: %w", err))
		}

		res := cfg.Timezone

		return &DailyQuestions{
			Timezone:    res,
			Deadlines:   make(map[string]int64),
			Responders:  make([]string, 0),
			muDeadlines: sync.RWMutex{},
			muResponses: sync.RWMutex{},
		}

	}).(*DailyQuestions)
}

func (q *DailyQuestions) GetDeadline(key string) *int64 {
	q.muDeadlines.RLock()
	defer q.muDeadlines.RUnlock()

	if q.Deadlines == nil {
		return nil
	}
	value, exists := q.Deadlines[key]
	if !exists {
		return nil
	}
	return &value
}

func (q *DailyQuestions) SetDeadlineTimer(eve *data.Event, nullifier string, deadline int64) {
	now := time.Now().UTC()

	q.muDeadlines.Lock()
	q.Deadlines[nullifier] = deadline
	q.muDeadlines.Unlock()

	time.AfterFunc(time.Duration(deadline)*time.Second, func() {
		q.muDeadlines.Lock()
		defer q.muDeadlines.Unlock()

		getTime := q.GetDeadline(nullifier)
		if getTime != nil && now.Unix()+deadline <= time.Now().UTC().Unix() {
			if eve != nil {
				delete(q.Deadlines, nullifier)
			}
		}
	})
}

func (q *DailyQuestions) ResponderExists(responder string) bool {
	q.muResponses.RLock()
	defer q.muResponses.RUnlock()

	for _, r := range q.Responders {
		if r == responder {
			return true
		}
	}
	return false
}

func (q *DailyQuestions) SetResponsesTimer(responder string, interval time.Duration) {
	q.muResponses.Lock()
	defer q.muResponses.Unlock()

	for _, r := range q.Responders {
		if r == responder {
		}
	}
	q.Responders = append(q.Responders, responder)

	time.AfterFunc(interval, func() {
		q.muResponses.Lock()
		defer q.muResponses.Unlock()

		for i, r := range q.Responders {
			if r == responder {
				q.Responders = append(q.Responders[:i], q.Responders[i+1:]...)
				break
			}
		}
	})
}

//func (q *DailyQuestions) RemoveAllQuestionsAtEndDay() {
//	now := time.Now().UTC()
//	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
//	timeUntilEndOfDay := endOfDay.Sub(now)
//
//	go func() {
//		time.Sleep(timeUntilEndOfDay)
//
//		q.muDeadlines.Lock()
//		defer q.muDeadlines.Unlock()
//
//		for nullifier := range q.Deadlines {
//			delete(q.Deadlines, nullifier)
//		}
//	}()
//}

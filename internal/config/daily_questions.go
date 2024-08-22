package config

import (
	"fmt"
	"sync"
	"time"

	"github.com/rarimo/geo-points-svc/internal/data"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
)

type DailyQuestions struct {
	Timezone    int
	Deadlines   map[string]int64
	Responders  map[string]struct{}
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
			Responders:  make(map[string]struct{}),
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

func (q *DailyQuestions) SetDeadlineTimer(log *logan.Entry, question data.DailyQuestionsQ, eve *data.Event, nullifier string, deadline int64) {
	now := time.Now().UTC()

	q.muDeadlines.Lock()
	q.Deadlines[nullifier] = deadline
	q.muDeadlines.Unlock()

	time.AfterFunc(time.Duration(deadline-now.Unix())*time.Second, func() {
		q.muDeadlines.Lock()
		defer q.muDeadlines.Unlock()

		if deadline <= time.Now().UTC().Unix() {
			if eve != nil {
				delete(q.Deadlines, nullifier)
			}
		}
	})

	err := question.IncrementAllParticipants()
	if err != nil {
		log.Infof("Failed to increment all participants: %v", err)
	}
}

func (q *DailyQuestions) ResponderExists(responder string) bool {
	q.muResponses.RLock()
	defer q.muResponses.RUnlock()

	_, exists := q.Responders[responder]
	return exists
}

func (q *DailyQuestions) SetResponsesTimer(responder string, interval time.Duration) {
	q.muResponses.Lock()

	if _, exists := q.Responders[responder]; exists {
		q.muResponses.Unlock()
		return
	}

	q.Responders[responder] = struct{}{}
	q.muResponses.Unlock()

	time.AfterFunc(interval, func() {
		q.muResponses.Lock()
		defer q.muResponses.Unlock()

		time.AfterFunc(interval, func() {
			q.muResponses.Lock()
			defer q.muResponses.Unlock()

			delete(q.Responders, responder)
		})
	})
}

func (q *DailyQuestions) ClearDeadlines() int {
	q.muDeadlines.Lock()
	defer q.muDeadlines.Unlock()

	count := len(q.Deadlines)
	q.Deadlines = make(map[string]int64)
	return count
}

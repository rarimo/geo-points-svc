package config

import (
	"fmt"
	"sync"
	"time"

	"github.com/rarimo/geo-points-svc/internal/data"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
)

var mu sync.Mutex
var endOfDayMutex sync.Mutex

type DailyQuestions struct {
	Timezone       int
	QuestionsQueue map[string]int64
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
			Timezone:       res,
			QuestionsQueue: make(map[string]int64),
		}

	}).(*DailyQuestions)
}

func (c DailyQuestions) InsertInQuestionsQueue(key string, value int64) {
	if c.QuestionsQueue == nil {
		c.QuestionsQueue = make(map[string]int64)
	}
	(c).QuestionsQueue[key] = value
}

func (c DailyQuestions) GetFromQuestionsQueue(key string) *int64 {
	if c.QuestionsQueue == nil {
		return nil
	}
	value, exists := c.QuestionsQueue[key]
	if !exists {
		return nil
	}
	return &value
}

func (c DailyQuestions) SetDailyQuestionTimeWithExpiration(eve *data.Event, nullifier string, deadline int64) {
	c.InsertInQuestionsQueue(nullifier, deadline)

	now := time.Now().UTC()

	go func() {
		time.Sleep(time.Duration(deadline) * time.Second)

		mu.Lock()
		defer mu.Unlock()

		getTime := c.GetFromQuestionsQueue(nullifier)
		if now.Unix() < *getTime+deadline {
			if eve != nil {
				delete(c.QuestionsQueue, nullifier)
			}
		}
	}()

	go func() {
		endOfDayMutex.Lock()
		defer endOfDayMutex.Unlock()

		c.RemoveAllQuestionsAtEndDay()
	}()
}

func (c DailyQuestions) RemoveAllQuestionsAtEndDay() {
	now := time.Now().UTC()
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	timeUntilEndOfDay := endOfDay.Sub(now)

	go func() {
		time.Sleep(timeUntilEndOfDay)

		mu.Lock()
		defer mu.Unlock()

		for nullifier := range c.QuestionsQueue {
			delete(c.QuestionsQueue, nullifier)
		}
	}()
}

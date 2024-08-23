package config

import (
	"fmt"
	"sync"
	"time"

	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
)

type DailyQuestions struct {
	Location *time.Location

	deadlines    map[string]Deadline
	lastDeadline *time.Time
	disabled     bool

	mu sync.Mutex
}

type Deadline struct {
	ID int
	At time.Time
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

		if cfg.Timezone < -12 || cfg.Timezone > 12 {
			panic(fmt.Errorf("timezone must be between -12 and 12"))
		}

		location := time.FixedZone(fmt.Sprint(cfg.Timezone), cfg.Timezone*int(time.Hour))

		return &DailyQuestions{
			Location:  location,
			deadlines: make(map[string]Deadline),
			mu:        sync.Mutex{},
		}

	}).(*DailyQuestions)
}

func (q *DailyQuestions) LocalTime(date time.Time) time.Time {
	return date.In(q.Location)
}

func (q *DailyQuestions) GetDeadline(nullifier string) *Deadline {
	q.mu.Lock()
	defer q.mu.Unlock()

	deadline, ok := q.deadlines[nullifier]
	if !ok {
		return nil
	}
	return &deadline
}

func (q *DailyQuestions) SetDeadline(nullifier string, id int, duration time.Duration) bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.disabled {
		return false
	}

	date := time.Now().UTC().Add(duration)
	q.deadlines[nullifier] = Deadline{ID: id, At: date}
	q.lastDeadline = &date
	return true
}

func (q *DailyQuestions) ClearDeadlines() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.disabled = true
	if q.lastDeadline != nil && q.lastDeadline.After(time.Now().UTC()) {
		q.mu.Unlock()
		timer := time.NewTimer(q.lastDeadline.Sub(time.Now().UTC()))
		<-timer.C
		q.mu.Lock()
	}
	q.disabled = false

	count := len(q.deadlines)
	q.deadlines = make(map[string]Deadline)
	return count
}

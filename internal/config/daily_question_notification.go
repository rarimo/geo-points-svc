package config

import (
	"fmt"

	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
)

type DailyQuestionsNotification struct {
	Body  string `fig:"body"`
	Title string `fig:"title"`
}

func (c *config) DailyQuestionsNotification() *DailyQuestionsNotification {
	return c.DailyQuestionNotification.Do(func() interface{} {
		var cfg DailyQuestionsNotification

		err := figure.Out(&cfg).
			From(kv.MustGetStringMap(c.getter, "daily_questions")).
			Please()
		if err != nil {
			panic(fmt.Errorf("failed to figure out daily questions config: %w", err))
		}

		return &DailyQuestionsNotification{
			Body:  cfg.Body,
			Title: cfg.Title,
		}

	}).(*DailyQuestionsNotification)
}

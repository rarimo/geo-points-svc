package config

type DailyQuestionTimeHash interface {
	SetDailyQuestionsTimeHash(key string, value int64)
	GetDailyQuestionsTimeHash() DailyQuestionsTimeHash
}

type DailyQuestionsTimeHash map[string]int64

func (c DailyQuestionsTimeHash) SetDailyQuestionsTimeHash(key string, value int64) {
	if c == nil {
		c = make(DailyQuestionsTimeHash)
	}
	(c)[key] = value
}

func (c DailyQuestionsTimeHash) GetDailyQuestionsTimeHash(key string) *int64 {
	if c == nil {
		return nil
	}
	value, exists := c[key]
	if !exists {
		return nil
	}
	return &value
}

func (c *config) DailyQuestionsTimeHash() DailyQuestionsTimeHash {
	return c.dailyQuestionsTimeHash
}

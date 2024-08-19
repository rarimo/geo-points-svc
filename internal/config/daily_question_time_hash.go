package config

type DailyQuestionsTimeHash interface {
	GetDailyQuestionsTimeHash() map[string]int64
	SetDailyQuestionsTimeHash(key string, value int64)
}

func (c *config) GetDailyQuestionsTimeHash() map[string]int64 {
	return c.dailyQuestionsTimeHash
}

func (c *config) SetDailyQuestionsTimeHash(key string, value int64) {
	c.dailyQuestionsTimeHash[key] = value
}

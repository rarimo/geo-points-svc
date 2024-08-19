package config

type DailyQuestionTimeHash interface {
	SetDailyQuestionsTimeHash(key string, value int64)
	GetDailyQuestionsTimeHash() DailyQuestionsTimeHash
}

type DailyQuestionTimeInfo struct {
	MaxDateToAnswer int64
	Answered        bool
}

type DailyQuestionsTimeHash map[string]DailyQuestionTimeInfo

func (c DailyQuestionsTimeHash) SetDailyQuestionsTimeHash(key string, value DailyQuestionTimeInfo) {
	if c == nil {
		c = make(DailyQuestionsTimeHash)
	}
	(c)[key] = value
}

func (c DailyQuestionsTimeHash) GetDailyQuestionsTimeHash(key string) *DailyQuestionTimeInfo {
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

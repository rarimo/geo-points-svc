/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import (
	"time"
)

type DailyQuestionAttributes struct {
	AnswerOptions map[string]interface{} `json:"answer_options"`
	// Event add date
	CreatedAt time.Time `json:"created_at"`
	// Reward for the right answer
	Reward int `json:"reward"`
	// Event start date, after which the event becomes active
	StartsAt int64 `json:"starts_at"`
	// Time for answer
	TimeForAnswer int64 `json:"time_for_answer"`
	// Title question
	Title string `json:"title"`
}

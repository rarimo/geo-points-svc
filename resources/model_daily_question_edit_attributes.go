/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "time"

type DailyQuestionEditAttributes struct {
	// right answer index
	CorrectAnswer int64 `json:"correct_answer"`
	// Answer options. Minimum 2, maximum 6
	Options []DailyQuestionOptions `json:"options"`
	// reward for a correct answer
	Reward int `json:"reward"`
	// start date when this question is available, hours and minutes are always 0
	StartsAt time.Time `json:"starts_at"`
	// time for answer
	TimeForAnswer int64 `json:"time_for_answer"`
	// Question title
	Title string `json:"title"`
}

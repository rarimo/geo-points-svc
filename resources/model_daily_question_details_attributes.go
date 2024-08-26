/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "time"

type DailyQuestionDetailsAttributes struct {
	// right answer index
	CorrectAnswer int64 `json:"correct_answer"`
	// start date when this question was create
	CreatedAt time.Time `json:"created_at"`
	// users who received the question, those who answered and those who did not answer in the time given to them
	NumAllParticipants int64 `json:"num_all_participants"`
	// number of correct answers
	NumCorrectAnswers int64 `json:"num_correct_answers"`
	// number of incorrect answers
	NumIncorrectAnswers int64 `json:"num_incorrect_answers"`
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

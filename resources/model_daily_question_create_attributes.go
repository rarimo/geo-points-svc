/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type DailyQuestionCreateAttributes struct {
	// Correct answer ID
	CorrectAnswer int64 `json:"correct_answer"`
	// Answer options. Minimum 2, maximum 6
	Options []DailyQuestionOptions `json:"options"`
	// Reward for a correct answer
	Reward int64 `json:"reward"`
	// Start date when this question is available, hours and minutes are always 0
	StartsAt string `json:"starts_at"`
	// Time for answer
	TimeForAnswer int64 `json:"time_for_answer"`
	// Question title
	Title string `json:"title"`
}

/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type DailyQuestionDetailsAttributes struct {
	// Correct answer ID
	CorrectAnswer int64 `json:"correct_answer"`
	// Start date when this question was create
	CreatedAt string `json:"created_at"`
	// Users who received the question, those who answered and those who did not answer in the time given to them
	NumAllParticipants int64 `json:"num_all_participants"`
	// Number of correct answers
	NumCorrectAnswers int64 `json:"num_correct_answers"`
	// Number of incorrect answers
	NumIncorrectAnswers int64 `json:"num_incorrect_answers"`
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

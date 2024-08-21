/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type DailyQuestionsStatusAttributes struct {
	// Time when the next question will be available.  If the time is in the past, then there is a question  on this day and the user has not yet answered it.  If the time is in the future, then the user has either  already answered the question on the current day or  there was no question on the current day.
	NextQuestionDate int64 `json:"next_question_date"`
	// The number of points the user will receive if they answer the question correctly.
	Reward int64 `json:"reward"`
	// The time within which the user has to answer this question after receiving it.
	TimeForAnswer int64 `json:"time_for_answer"`
}

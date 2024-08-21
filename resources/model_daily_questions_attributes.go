/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type DailyQuestionsAttributes struct {
	// Time limit after which it is impossible to answer the question. Calculated as current time + time for answer
	Deadline int64 `json:"deadline"`
	// Answer options. Minimum 2, maximum 6
	Options []DailyQuestionOptions `json:"options"`
	// Question title
	Title string `json:"title"`
}

/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type DailyQuestionAttributes struct {
	// Status activity event
	Active *bool `json:"Active,omitempty"`
	// ID question
	ID            int                    `json:"ID"`
	AnswerOptions map[string]interface{} `json:"answer_options"`
	// Event add date
	CreatedAt int `json:"created_at"`
	// Reward for the right answer
	Reward int `json:"reward"`
	// Event start date, after which the event becomes active
	StartsAt int `json:"starts_at"`
	// Time for answer
	TimeForAnswer int `json:"time_for_answer"`
	// Title question
	Title string `json:"title"`
}

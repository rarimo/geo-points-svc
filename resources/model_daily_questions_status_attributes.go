/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type DailyQuestionsStatusAttributes struct {
	// question of the day already done for user
	AlreadyDoneForUser bool `json:"already_done_for_user"`
	// availability questions of the day
	Availability bool `json:"availability"`
	// time to next question of the day
	TimeToNext int64 `json:"time_to_next"`
}

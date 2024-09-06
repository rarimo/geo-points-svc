package referralid

import "time"

func CheckOpportunityChange(now time.Time, timeReq time.Time, location *time.Location) bool {
	now = now.AddDate(0, 0, -1)
	if !timeReq.After(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, location)) {
		return false
	}
	return true
}

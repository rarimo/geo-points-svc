package config

import "time"

type TimeZoneProvider interface {
	UtcTimeZone() *time.Location
}

func (c *config) UtcTimeZone() *time.Location {
	return c.timeZone
}

func NowTime() time.Time {
	loc, _ := time.LoadLocation(locationName)
	return time.Now().In(loc)
}

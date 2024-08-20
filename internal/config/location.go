package config

import "time"

type Location = *time.Location

func (c *config) Location() *time.Location {
	return c.timeLocation
}

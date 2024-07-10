package models

import (
	"time"

	"github.com/rarimo/geo-points-svc/resources"
)

type EventType struct {
	Name             string     `fig:"name,required" db:"name"`
	Description      string     `fig:"description,required" db:"description"`
	ShortDescription string     `fig:"short_description,required" db:"short_description"`
	Reward           int64      `fig:"reward,required" db:"reward"`
	Title            string     `fig:"title,required" db:"title"`
	Frequency        Frequency  `fig:"frequency,required" db:"frequency"`
	StartsAt         *time.Time `fig:"starts_at" db:"starts_at"`
	ExpiresAt        *time.Time `fig:"expires_at" db:"expires_at"`
	NoAutoOpen       bool       `fig:"no_auto_open" db:"no_auto_open"`
	AutoClaim        bool       `fig:"auto_claim" db:"auto_claim"`
	Disabled         bool       `fig:"disabled" db:"disabled"`
	ActionURL        *string    `fig:"action_url" db:"action_url"`
	Logo             *string    `fig:"logo" db:"logo"`
	QRCodeValue      *string    `fig:"qr_code_value" db:"qr_code_value"`
}

func ResourceToModel(r resources.EventStaticMeta) EventType {
	// intended that no_auto_open field is not accessible through API due to being
	// related only to back-end
	return EventType{
		Name:             r.Name,
		Description:      r.Description,
		ShortDescription: r.ShortDescription,
		Reward:           r.Reward,
		Title:            r.Title,
		Frequency:        Frequency(r.Frequency),
		StartsAt:         r.StartsAt,
		ExpiresAt:        r.ExpiresAt,
		AutoClaim:        r.AutoClaim,
		Disabled:         r.Disabled,
		ActionURL:        r.ActionUrl,
		Logo:             r.Logo,
		QRCodeValue:      r.QrCodeValue,
	}
}

func (e EventType) Flag() string {
	switch {
	case e.Disabled:
		return FlagDisabled
	case e.StartsAt != nil && e.StartsAt.After(time.Now().UTC()):
		return FlagNotStarted
	case e.ExpiresAt != nil && e.ExpiresAt.Before(time.Now().UTC()):
		return FlagExpired
	default:
		return FlagActive
	}
}

func (e EventType) Resource(qr bool) resources.EventStaticMeta {

	res := resources.EventStaticMeta{
		Name:             e.Name,
		Description:      e.Description,
		ShortDescription: e.ShortDescription,
		Reward:           e.Reward,
		Title:            e.Title,
		Frequency:        e.Frequency.String(),
		StartsAt:         e.StartsAt,
		ExpiresAt:        e.ExpiresAt,
		AutoClaim:        e.AutoClaim,
		ActionUrl:        e.ActionURL,
		Disabled:         e.Disabled,
		Logo:             e.Logo,
		Flag:             e.Flag(),
	}

	if qr {
		res.QrCodeValue = e.QRCodeValue
	}

	return res
}

func (e EventType) ForUpdate() map[string]any {
	return map[string]any{
		"description":       e.Description,
		"short_description": e.ShortDescription,
		"reward":            e.Reward,
		"title":             e.Title,
		"frequency":         e.Frequency,
		"starts_at":         e.StartsAt,
		"expires_at":        e.ExpiresAt,
		"auto_claim":        e.AutoClaim,
		"disabled":          e.Disabled,
		"action_url":        e.ActionURL,
		"logo":              e.Logo,
		"qr_code_value":     e.QRCodeValue,
	}
}

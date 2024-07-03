package models

import (
	"net/url"
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
	ActionURL        *url.URL   `fig:"action_url" db:"action_url"`
	Logo             *url.URL   `fig:"logo" db:"logo"`
	QRCodeValue      *string    `fig:"qr_code_value" db:"qr_code_value"`
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

func (e EventType) Resource() resources.EventStaticMeta {
	safeConv := func(u *url.URL) *string {
		if u == nil {
			return nil
		}
		s := u.String()
		return &s
	}

	return resources.EventStaticMeta{
		Name:             e.Name,
		Description:      e.Description,
		ShortDescription: e.ShortDescription,
		Reward:           e.Reward,
		Title:            e.Title,
		Frequency:        e.Frequency.String(),
		StartsAt:         e.StartsAt,
		ExpiresAt:        e.ExpiresAt,
		ActionUrl:        safeConv(e.ActionURL),
		Logo:             safeConv(e.Logo),
		Flag:             e.Flag(),
		QrCodeValue:      e.QRCodeValue,
	}
}

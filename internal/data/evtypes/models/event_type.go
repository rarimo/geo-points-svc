package models

import (
	"strings"
	"time"

	"github.com/rarimo/geo-points-svc/resources"
)

type EventType struct {
	Name             string          `fig:"name,required" db:"name"`
	Reward           int64           `fig:"reward,required" db:"reward"`
	Title            string          `fig:"title,required" db:"title"`
	Description      string          `fig:"description,required" db:"description"`
	ShortDescription string          `fig:"short_description,required" db:"short_description"`
	Localized        LocalizationMap `fig:"localized" db:"localized"`
	Frequency        Frequency       `fig:"frequency,required" db:"frequency"`
	StartsAt         *time.Time      `fig:"starts_at" db:"starts_at"`
	ExpiresAt        *time.Time      `fig:"expires_at" db:"expires_at"`
	NoAutoOpen       bool            `fig:"no_auto_open" db:"no_auto_open"`
	AutoClaim        bool            `fig:"auto_claim" db:"auto_claim"`
	Disabled         bool            `fig:"disabled" db:"disabled"`
	ActionURL        *string         `fig:"action_url" db:"action_url"`
	Logo             *string         `fig:"logo" db:"logo"`
	QRCodeValue      *string         `db:"qr_code_value"`
	PollEventID      *string         `db:"poll_event_id"`
	PollContract     *string         `db:"poll_contract"`
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
		PollEventID:      r.PollEventId,
		PollContract:     r.PollContract,
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

func (e EventType) Resource(locale string) resources.EventStaticMeta {
	l := e.GetLocalized(strings.ToLower(locale))
	return resources.EventStaticMeta{
		Name:             e.Name,
		Reward:           e.Reward,
		Description:      l.Description,
		ShortDescription: l.ShortDescription,
		Title:            l.Title,
		Frequency:        e.Frequency.String(),
		StartsAt:         e.StartsAt,
		ExpiresAt:        e.ExpiresAt,
		AutoClaim:        e.AutoClaim,
		ActionUrl:        e.ActionURL,
		Disabled:         e.Disabled,
		Logo:             e.Logo,
		Flag:             e.Flag(),
	}
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
	}
}

func (e EventType) GetLocalized(locale string) Localized {
	def := Localized{
		Title:            e.Title,
		Description:      e.Description,
		ShortDescription: e.ShortDescription,
	}

	if len(e.Localized) == 0 {
		return def
	}

	v, ok := e.Localized[locale]
	if !ok {
		return def
	}

	return v
}

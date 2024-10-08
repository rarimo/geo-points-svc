package models

type Frequency string

func (f Frequency) String() string {
	return string(f)
}

const (
	OneTime   Frequency = "one-time"
	Daily     Frequency = "daily"
	Weekly    Frequency = "weekly"
	Unlimited Frequency = "unlimited"
)

const (
	TypeFreeWeekly           = "free_weekly"
	TypeBeReferred           = "be_referred"
	TypeReferralSpecific     = "referral_specific"
	TypePassportScan         = "passport_scan"
	TypeInternalPassportScan = "internal_passport_scan"
	TypeExternalPassportScan = "external_passport_scan"
	TypePollParticipation    = "poll_participation"
	TypeEarlyTest            = "early_test"
	TypeDailyQuestion        = "daily_question"
	TypeBonusCode            = "bonus_code"
)

const (
	FlagActive     = "active"
	FlagNotStarted = "not_started"
	FlagExpired    = "expired"
	FlagDisabled   = "disabled"
)

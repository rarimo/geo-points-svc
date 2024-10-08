package handlers

import (
	"context"
	"net/http"

	"github.com/rarimo/geo-auth-svc/pkg/hmacsig"
	"github.com/rarimo/geo-auth-svc/resources"
	"github.com/rarimo/geo-points-svc/internal/config"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes"
	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	eventsQCtxKey
	balancesQCtxKey
	referralsQCtxKey
	eventTypesCtxKey
	eventTypesQCtxKey
	withdrawalsQCtxKey
	userClaimsCtxKey
	levelsCtxKey
	verifiersCtxKey
	sigCalculatorCtxKey
	voteVerifierCtxKey
	dailyQuestionsCtxKey
	dailyQuestionsCfgCtxKey
	abstractionCtxKey
	bonusCodesQCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxEventsQ(q data.EventsQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, eventsQCtxKey, q)
	}
}

func EventsQ(r *http.Request) data.EventsQ {
	return r.Context().Value(eventsQCtxKey).(data.EventsQ).New()
}

func CtxBalancesQ(q data.BalancesQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, balancesQCtxKey, q)
	}
}

func CtxDailyQuestionsQ(q data.DailyQuestionsQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, dailyQuestionsCtxKey, q)
	}
}

func BalancesQ(r *http.Request) data.BalancesQ {
	return r.Context().Value(balancesQCtxKey).(data.BalancesQ).New()
}

func DailyQuestionsQ(r *http.Request) data.DailyQuestionsQ {
	return r.Context().Value(dailyQuestionsCtxKey).(data.DailyQuestionsQ).New()
}

func CtxReferralsQ(q data.ReferralsQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, referralsQCtxKey, q)
	}
}

func ReferralsQ(r *http.Request) data.ReferralsQ {
	return r.Context().Value(referralsQCtxKey).(data.ReferralsQ).New()
}

func CtxEventTypes(types *evtypes.Types) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, eventTypesCtxKey, types)
	}
}

func EventTypes(r *http.Request) *evtypes.Types {
	return r.Context().Value(eventTypesCtxKey).(*evtypes.Types)
}

func CtxEventTypesQ(q data.EventTypesQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, eventTypesQCtxKey, q)
	}
}

func EventTypesQ(r *http.Request) data.EventTypesQ {
	return r.Context().Value(eventTypesQCtxKey).(data.EventTypesQ).New()
}

func CtxWithdrawalsQ(q data.WithdrawalsQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, withdrawalsQCtxKey, q)
	}
}

func WithdrawalsQ(r *http.Request) data.WithdrawalsQ {
	return r.Context().Value(withdrawalsQCtxKey).(data.WithdrawalsQ).New()
}

func CtxBonusCodesQ(q data.BonusCodesQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, bonusCodesQCtxKey, q)
	}
}

func BonusCodesQ(r *http.Request) data.BonusCodesQ {
	return r.Context().Value(bonusCodesQCtxKey).(data.BonusCodesQ).New()
}

func CtxUserClaims(claim []resources.Claim) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, userClaimsCtxKey, claim)
	}
}

func UserClaims(r *http.Request) []resources.Claim {
	return r.Context().Value(userClaimsCtxKey).([]resources.Claim)
}

func CtxVerifiers(v config.Verifiers) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, verifiersCtxKey, v)
	}
}

func Verifiers(r *http.Request) config.Verifiers {
	return r.Context().Value(verifiersCtxKey).(config.Verifiers)
}

func CtxDailyQuestion(v *config.DailyQuestions) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, dailyQuestionsCfgCtxKey, v)
	}
}

func DailyQuestions(r *http.Request) *config.DailyQuestions {
	return r.Context().Value(dailyQuestionsCfgCtxKey).(*config.DailyQuestions)
}

func CtxPollVerifier(v *config.PollVerifier) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, voteVerifierCtxKey, v)
	}
}

func PollVerifier(r *http.Request) *config.PollVerifier {
	return r.Context().Value(voteVerifierCtxKey).(*config.PollVerifier)
}

func CtxSigCalculator(calc hmacsig.Calculator) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, sigCalculatorCtxKey, calc)
	}
}

func SigCalculator(r *http.Request) hmacsig.Calculator {
	return r.Context().Value(sigCalculatorCtxKey).(hmacsig.Calculator)
}

func CtxLevels(levels *config.Levels) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, levelsCtxKey, levels)
	}
}

func Levels(r *http.Request) *config.Levels {
	return r.Context().Value(levelsCtxKey).(*config.Levels)
}

func CtxAbstraction(abstraction *config.AbstractionConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, abstractionCtxKey, abstraction)
	}
}

func Abstraction(r *http.Request) *config.AbstractionConfig {
	return r.Context().Value(abstractionCtxKey).(*config.AbstractionConfig)
}

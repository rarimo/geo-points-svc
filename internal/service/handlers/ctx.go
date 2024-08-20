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
	userClaimsCtxKey
	levelsCtxKey
	verifiersCtxKey
	sigCalculatorCtxKey
	voteVerifierCtxKey
	dailyQuestionsCtxKey
	dailyQuestionsHashCtxKey
	locationCtxKey
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

func CtxLocation(v config.Location) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, locationCtxKey, v)
	}
}

func Location(r *http.Request) config.Location {
	return r.Context().Value(locationCtxKey).(config.Location)
}

func CtxDailyQuestionTimeHash(v config.DailyQuestionsTimeHash) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, dailyQuestionsHashCtxKey, v)
	}
}

func DailyQuestionTimeHash(r *http.Request) config.DailyQuestionsTimeHash {
	return r.Context().Value(dailyQuestionsHashCtxKey).(config.DailyQuestionsTimeHash)
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

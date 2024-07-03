package handlers

import (
	"context"
	"net/http"

	"github.com/rarimo/decentralized-auth-svc/resources"
	"github.com/rarimo/geo-points-svc/internal/config"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes"
	zk "github.com/rarimo/zkverifier-kit"
	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	eventsQCtxKey
	balancesQCtxKey
	referralsQCtxKey
	eventTypesCtxKey
	userClaimsCtxKey
	levelsCtxKey
	verifierCtxKey
	sigVerifierCtxKey
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

func BalancesQ(r *http.Request) data.BalancesQ {
	return r.Context().Value(balancesQCtxKey).(data.BalancesQ).New()
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

func CtxUserClaims(claim []resources.Claim) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, userClaimsCtxKey, claim)
	}
}

func UserClaims(r *http.Request) []resources.Claim {
	return r.Context().Value(userClaimsCtxKey).([]resources.Claim)
}

func CtxVerifier(verifier *zk.Verifier) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, verifierCtxKey, verifier)
	}
}

func Verifier(r *http.Request) *zk.Verifier {
	return r.Context().Value(verifierCtxKey).(*zk.Verifier)
}

func CtxSigVerifier(sigVerifier []byte) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, sigVerifierCtxKey, sigVerifier)
	}
}

func SigVerifier(r *http.Request) []byte {
	return r.Context().Value(sigVerifierCtxKey).([]byte)
}

func CtxLevels(levels config.Levels) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, levelsCtxKey, levels)
	}
}

func Levels(r *http.Request) config.Levels {
	return r.Context().Value(levelsCtxKey).(config.Levels)
}

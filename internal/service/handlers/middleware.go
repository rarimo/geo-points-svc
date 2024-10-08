package handlers

import (
	"context"
	"net/http"

	"github.com/rarimo/geo-auth-svc/pkg/auth"
	"github.com/rarimo/geo-points-svc/internal/data/pg"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3"
)

func AuthMiddleware(auth *auth.Client, log *logan.Entry) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, err := auth.ValidateJWT(r)
			if err != nil {
				log.WithError(err).Info("Got invalid auth or validation error")
				ape.RenderErr(w, problems.Unauthorized())
				return
			}

			if len(claims) == 0 {
				log.Debug("Claims are empty")
				ape.RenderErr(w, problems.Unauthorized())
				return
			}

			ctx := CtxUserClaims(claims)(r.Context())
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// DBCloneMiddleware is designed to clone DB session on each request. You must
// put all new DB handlers here instead of ape.CtxMiddleware, unless you have a
// reason to do otherwise.
func DBCloneMiddleware(db *pgdb.DB) func(http.Handler) http.Handler {
	type ctxExtender func(context.Context) context.Context

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			clone := db.Clone()
			ctx := r.Context()

			extenders := []ctxExtender{
				CtxEventsQ(pg.NewEvents(clone)),
				CtxBalancesQ(pg.NewBalances(clone)),
				CtxReferralsQ(pg.NewReferrals(clone)),
				CtxEventTypesQ(pg.NewEventTypes(clone)),
				CtxDailyQuestionsQ(pg.NewDailyQuestionsQ(clone)),
				CtxWithdrawalsQ(pg.NewWithdrawals(clone)),
				CtxBonusCodesQ(pg.NewBonusCodesQ(clone)),
			}

			for _, extender := range extenders {
				ctx = extender(ctx)
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

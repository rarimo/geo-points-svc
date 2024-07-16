package service

import (
	"context"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/go-chi/chi"
	"github.com/rarimo/geo-points-svc/internal/config"
	"github.com/rarimo/geo-points-svc/internal/service/handlers"
	"gitlab.com/distributed_lab/ape"
)

func Run(ctx context.Context, cfg config.Config) {
	setBech32Prefixes()
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(cfg.Log()),
		ape.LoganMiddleware(cfg.Log()),
		ape.CtxMiddleware(
			handlers.CtxLog(cfg.Log()),
			handlers.CtxEventTypes(cfg.EventTypes()),
			handlers.CtxLevels(cfg.Levels()),
			handlers.CtxVerifier(cfg.Verifier()),
			handlers.CtxSigCalculator(cfg.SigCalculator()),
		),
		handlers.DBCloneMiddleware(cfg.DB()),
	)

	authMW := handlers.AuthMiddleware(cfg.Auth(), cfg.Log())
	r.Route("/integrations/geo-points-svc/v1", func(r chi.Router) {
		r.Route("/public", func(r chi.Router) {
			r.Route("/balances", func(r chi.Router) {
				r.Use(authMW)
				r.Post("/", handlers.CreateBalance)
				r.Route("/{nullifier}", func(r chi.Router) {
					r.Get("/", handlers.GetBalance)
					r.Post("/verifypassport", handlers.VerifyPassport)
					r.Post("/join_program", handlers.VerifyPassport)
				})
			})
			r.Route("/events", func(r chi.Router) {
				r.Use(authMW)
				r.Get("/", handlers.ListEvents)
				r.Get("/{id}", handlers.GetEvent)
				r.Patch("/{id}/qrcode", handlers.FulfillQREvent)
				r.Patch("/{id}", handlers.ClaimEvent)
			})
			r.Get("/balances", handlers.Leaderboard)
			r.Route("/event_types", func(r chi.Router) {
				r.With(authMW).Get("/qr", handlers.ListQREventTypes)
				r.Get("/", handlers.ListEventTypes)
				r.With(authMW).Post("/", handlers.CreateEventType)
				r.Get("/{name}", handlers.GetEventType)
				r.With(authMW).Put("/{name}", handlers.UpdateEventType)
			})
		})
		// must be accessible only within the cluster
		r.Route("/private", func(r chi.Router) {
			r.Post("/referrals", handlers.EditReferrals)
		})
	})

	r.Route("/integrations/geo-points-svc/v2/public/balances", func(r chi.Router) {
		r.Use(authMW)
		r.Post("/", handlers.CreateBalanceV2)
		r.Patch("/{nullifier}", handlers.ActivateBalance)
		r.Post("/{nullifier}/verifypassport", handlers.VerifyPassportV2)
	})

	cfg.Log().Info("Service started")
	ape.Serve(ctx, r, cfg, ape.ServeOpts{})
}

func setBech32Prefixes() {
	c := types.GetConfig()
	c.SetBech32PrefixForAccount("rarimo", "rarimopub")
	c.SetBech32PrefixForValidator("rarimovaloper", "rarimovaloperpub")
	c.SetBech32PrefixForConsensusNode("rarimovalcons", "rarimovalconspub")
	c.Seal()
}

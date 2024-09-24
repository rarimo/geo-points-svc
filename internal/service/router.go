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
			handlers.CtxVerifiers(cfg.Verifiers()),
			handlers.CtxSigCalculator(cfg.SigCalculator()),
			handlers.CtxPollVerifier(cfg.PollVerifier()),
			handlers.CtxDailyQuestion(cfg.DailyQuestions()),
			handlers.CtxAbstraction(cfg.AbstractionConfig()),
		),
		handlers.DBCloneMiddleware(cfg.DB()),
	)

	authMW := handlers.AuthMiddleware(cfg.Auth(), cfg.Log())
	r.Route("/integrations/geo-points-svc/v1", func(r chi.Router) {
		r.Route("/public", func(r chi.Router) {
			r.Route("/abstraction", func(r chi.Router) {
				r.Route("/accounts", func(r chi.Router) {
					r.Use(authMW)
					r.Post("/", handlers.CreateAbstractionAccount)
					r.Get("/{nullifier}", handlers.GetAbstractionAccount)
				})
			})

			r.Route("/balances", func(r chi.Router) {
				r.Use(authMW)
				r.Post("/", handlers.CreateBalance)
				r.Route("/{nullifier}", func(r chi.Router) {
					r.Get("/", handlers.GetBalance)
					r.Patch("/", handlers.ActivateBalance)
					r.Post("/verifypassport", handlers.VerifyInternalPassport)
					r.Post("/join_program", handlers.VerifyInternalPassport)
					r.Route("/verify", func(r chi.Router) {
						r.Post("/external", handlers.VerifyExternalPassport)
					})
					r.Post("/withdrawals", handlers.Withdraw)
				})
			})

			r.Route("/daily_questions", func(r chi.Router) {
				r.Use(authMW)
				r.Route("/admin", func(r chi.Router) {
					r.Delete("/{question_id}", handlers.DeleteDailyQuestion)
					r.Patch("/{question_id}", handlers.EditDailyQuestion)
					r.Post("/", handlers.CreateDailyQuestion)
					r.Get("/", handlers.FilterStartAtDailyQuestions)
				})
				r.Route("/{nullifier}", func(r chi.Router) {
					r.Get("/status", handlers.GetDailyQuestionsStatus)
					r.Get("/", handlers.GetDailyQuestion)
					r.Post("/", handlers.CheckDailyQuestion)
				})
			})
			r.Route("/events", func(r chi.Router) {
				r.Use(authMW)
				r.Get("/", handlers.ListEvents)
				r.Post("/poll", handlers.FulfillPollEvent)
				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", handlers.GetEvent)
					r.Patch("/", handlers.ClaimEvent)
					r.Patch("/qrcode", handlers.FulfillQREvent)
				})
			})
			r.Get("/balances", handlers.Leaderboard)
			r.Route("/event_types", func(r chi.Router) {
				r.Get("/", handlers.ListEventTypes)
				r.Get("/{name}", handlers.GetEventType)
				r.With(authMW).Get("/qr", handlers.ListQREventTypes)
				r.With(authMW).Post("/", handlers.CreateEventType)
				r.With(authMW).Put("/{name}", handlers.UpdateEventType)
			})
		})
		// must be accessible only within the cluster
		r.Route("/private", func(r chi.Router) {
			r.Post("/referrals", handlers.EditReferrals)
		})
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

package config

import (
	"github.com/rarimo/geo-auth-svc/pkg/auth"
	"github.com/rarimo/geo-auth-svc/pkg/hmacsig"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes"
	"github.com/rarimo/saver-grpc-lib/broadcaster"
	"github.com/rarimo/zkverifier-kit/root"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type Config interface {
	comfig.Logger
	pgdb.Databaser
	comfig.Listenerer
	auth.Auther //nolint:misspell
	broadcaster.Broadcasterer
	evtypes.EventTypeser
	hmacsig.SigCalculatorProvider
	PollVerifierer
	Abstraction

	Levels() *Levels
	Verifiers() Verifiers
	DailyQuestions() *DailyQuestions
}

type config struct {
	comfig.Logger
	pgdb.Databaser
	comfig.Listenerer
	auth.Auther
	broadcaster.Broadcasterer
	evtypes.EventTypeser
	hmacsig.SigCalculatorProvider
	PollVerifierer
	Abstraction

	passport root.VerifierProvider

	DailyQuestion comfig.Once
	levels        comfig.Once
	verifier      comfig.Once
	getter        kv.Getter
}

func New(getter kv.Getter) Config {
	return &config{
		getter:                getter,
		Databaser:             pgdb.NewDatabaser(getter),
		Listenerer:            comfig.NewListenerer(getter),
		Logger:                comfig.NewLogger(getter, comfig.LoggerOpts{}),
		Auther:                auth.NewAuther(getter), //nolint:misspell
		PollVerifierer:        NewPollVerifier(getter),
		Broadcasterer:         broadcaster.New(getter),
		passport:              root.NewVerifierProvider(getter, root.PoseidonSMT),
		EventTypeser:          evtypes.NewConfig(getter),
		SigCalculatorProvider: hmacsig.NewCalculatorProvider(getter),
		Abstraction:           NewAbstractionConfig(getter),
	}
}

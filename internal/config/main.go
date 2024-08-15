package config

import (
	"log"
	"time"

	"github.com/rarimo/geo-auth-svc/pkg/auth"
	"github.com/rarimo/geo-auth-svc/pkg/hmacsig"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes"
	"github.com/rarimo/saver-grpc-lib/broadcaster"
	"github.com/rarimo/zkverifier-kit/root"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const locationName = "Asia/Tbilisi"

type Config interface {
	comfig.Logger
	pgdb.Databaser
	comfig.Listenerer
	auth.Auther //nolint:misspell
	broadcaster.Broadcasterer
	evtypes.EventTypeser
	hmacsig.SigCalculatorProvider
	PollVerifierer
	TimeZoneProvider

	Levels() *Levels
	Verifiers() Verifiers
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

	timeZone *time.Location

	passport root.VerifierProvider

	levels   comfig.Once
	verifier comfig.Once
	getter   kv.Getter
}

func New(getter kv.Getter) Config {
	location, err := time.LoadLocation(locationName)
	if err != nil {
		log.Fatalf("error loading location: %v", err)
	}
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
		timeZone:              location,
	}
}

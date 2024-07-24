package config

import (
	"fmt"

	zk "github.com/rarimo/zkverifier-kit"
	"github.com/rarimo/zkverifier-kit/root"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
)

const (
	proofEventIDValue       = "111186066134341633902189494613533900917417361106374681011849132651019822199"
	proofSelectorValue      = "236065"
	maxIdentityCount        = 1
	documentTypeID          = "ID"
	passportVerificationKey = "./proof_keys/passport.json"
	pollVerificationKey     = "./proof_keys/poll.json"
)

type Verifiers struct {
	Passport *zk.Verifier
	Poll     *zk.Verifier
	PollRoot *root.ProposalSMTVerifier
}

func (c *config) Verifiers() Verifiers {
	return c.verifier.Do(func() interface{} {
		var cfg struct {
			AllowedAge               int   `fig:"allowed_age,required"`
			AllowedIdentityTimestamp int64 `fig:"allowed_identity_timestamp,required"`
		}

		err := figure.
			Out(&cfg).
			From(kv.MustGetStringMap(c.getter, "verifier")).
			Please()
		if err != nil {
			panic(fmt.Errorf("failed to figure out verifier: %w", err))
		}

		pass, err := zk.NewVerifier(nil,
			zk.WithProofType(zk.GeorgianPassport),
			zk.WithCitizenships("GEO"),
			zk.WithVerificationKeyFile(passportVerificationKey),
			zk.WithAgeAbove(cfg.AllowedAge),
			zk.WithPassportRootVerifier(c.passport.ProvideVerifier()),
			zk.WithProofSelectorValue(proofSelectorValue),
			zk.WithEventID(proofEventIDValue),
			zk.WithIdentitiesCounter(maxIdentityCount),
			zk.WithIdentitiesCreationTimestampLimit(cfg.AllowedIdentityTimestamp),
			zk.WithDocumentType(documentTypeID),
		)
		if err != nil {
			panic(fmt.Errorf("failed to initialize passport verifier: %w", err))
		}

		poll, err := zk.NewVerifier(nil,
			zk.WithProofType(zk.PollParticipation),
			zk.WithEventID(proofEventIDValue),
			zk.WithVerificationKeyFile(pollVerificationKey))
		if err != nil {
			panic(fmt.Errorf("failed to initialize poll verifier: %w", err))
		}

		rv := c.poll.ProvideVerifier().(*root.ProposalSMTVerifier)

		return Verifiers{
			Passport: pass,
			Poll:     poll,
			PollRoot: rv,
		}
	}).(Verifiers)
}

package config

import (
	"fmt"

	zk "github.com/rarimo/zkverifier-kit"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
)

const (
	proofEventIDValue  = "111186066134341633902189494613533900917417361106374681011849132651019822199"
	proofSelectorValue = "236065"
	maxIdentityCount   = 1
	documentTypeID     = "ID"
)

func (c *config) Verifier() *zk.Verifier {
	return c.verifier.Do(func() interface{} {
		var cfg struct {
			AllowedAge               int    `fig:"allowed_age,required"`
			VerificationKeyPath      string `fig:"verification_key_path,required"`
			AllowedIdentityTimestamp int64  `fig:"allowed_identity_timestamp,required"`
		}

		err := figure.
			Out(&cfg).
			From(kv.MustGetStringMap(c.getter, "verifier")).
			Please()
		if err != nil {
			panic(fmt.Errorf("failed to figure out verifier: %w", err))
		}

		v, err := zk.NewVerifier(nil,
			zk.WithCitizenships("GEO"),
			zk.WithProofType(zk.GeorgianPassport),
			zk.WithVerificationKeyFile(cfg.VerificationKeyPath),
			zk.WithAgeAbove(cfg.AllowedAge),
			zk.WithIdentityVerifier(c.ProvideVerifier()),
			zk.WithProofSelectorValue(proofSelectorValue),
			zk.WithEventID(proofEventIDValue),
			zk.WithIdentitiesCounter(maxIdentityCount),
			zk.WithIdentitiesCreationTimestampLimit(cfg.AllowedIdentityTimestamp),
			zk.WithDocumentType(documentTypeID),
		)

		if err != nil {
			panic(fmt.Errorf("failed to initialize passport verifier: %w", err))
		}

		return v
	}).(*zk.Verifier)
}

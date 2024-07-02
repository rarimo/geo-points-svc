package config

import (
	"encoding/hex"
	"fmt"

	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
)

func (c *config) SigVerifier() []byte {
	return c.sigVerifier.Do(func() interface{} {
		var cfg struct {
			VerificationKey string `fig:"verification_key,required"`
		}

		err := figure.Out(&cfg).
			From(kv.MustGetStringMap(c.getter, "sig_verifier")).
			Please()
		if err != nil {
			panic(fmt.Errorf("failed to figure out sig_verifier: %w", err))
		}

		key, err := hex.DecodeString(cfg.VerificationKey)
		if err != nil {
			panic(fmt.Errorf("verification_key is not a hex: %w", err))
		}

		return key
	}).([]byte)
}

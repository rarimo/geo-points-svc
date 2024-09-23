package config

import (
	"fmt"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
)

type Creds struct {
	Path string `fig:"path"`
}

func (c *config) Creds() *Creds {
	return c.Cred.Do(func() interface{} {
		var cfg Creds

		err := figure.Out(&cfg).
			From(kv.MustGetStringMap(c.getter, "creds")).
			Please()
		if err != nil {
			panic(fmt.Errorf("failed to figure creds: %w", err))
		}

		return &cfg
	}).(*Creds)
}

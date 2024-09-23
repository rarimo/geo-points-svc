package config

import (
	"fmt"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type Creds interface {
	Creds() *CredsDetails
}

func NewCreds(getter kv.Getter) Creds {
	return &credsDetails{
		getter: getter,
	}
}

type credsDetails struct {
	once   comfig.Once
	getter kv.Getter
}

type CredsDetails struct {
	Path string `fig:"creds,required"`
}

func (c *credsDetails) Creds() *CredsDetails {
	return c.once.Do(func() interface{} {

		var cfg CredsDetails

		err := figure.Out(&cfg).
			From(kv.MustGetStringMap(c.getter, "creds")).
			Please()
		if err != nil {
			panic(fmt.Errorf("failed to figure creads: %w", err))
		}

		return &cfg

	}).(*CredsDetails)
}

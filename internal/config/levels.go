package config

import (
	"errors"
	"fmt"
	"slices"

	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
)

type Level struct {
	Level             int  `fig:"lvl,required"`
	Threshold         int  `fig:"threshold,required"`
	Referrals         int  `fig:"referrals"`
	Infinity          bool `fig:"infinity"`
	WithdrawalAllowed bool `fig:"withdrawal_allowed"`
}

type Levels struct {
	levels        map[int]Level
	Downgradeable bool
}

func (c *config) Levels() *Levels {
	return c.levels.Do(func() interface{} {
		var cfg struct {
			Downgradeable bool    `fig:"downgradeable"`
			Lvls          []Level `fig:"levels,required"`
		}

		err := figure.Out(&cfg).
			From(kv.MustGetStringMap(c.getter, "levels")).
			Please()
		if err != nil {
			panic(fmt.Errorf("failed to figure out levels config: %w", err))
		}

		if len(cfg.Lvls) == 0 {
			panic(errors.New("no levels provided in config"))
		}

		res := make(map[int]Level, len(cfg.Lvls))
		for _, v := range cfg.Lvls {
			res[v.Level] = v
		}

		return &Levels{
			levels:        res,
			Downgradeable: cfg.Downgradeable,
		}
	}).(*Levels)
}

func (l *Levels) WithdrawalAllowed(level int) bool {
	levelConfig, ok := l.levels[level]
	if !ok {
		return false
	}

	return levelConfig.WithdrawalAllowed
}

// LvlUp Calculates new lvl. New lvl always greater then current level
func (l *Levels) LvlChange(currentLevel int, totalAmount int64) (refCoundToAdd *int, newLevel int) {
	var downgrade bool
	if l.Downgradeable && l.levels[currentLevel].Threshold > int(totalAmount) {
		downgrade = true
	}
	lvls := make([]int, 0, len(l.levels))
	refCoundToAdd = new(int)

	if downgrade {
		for k, v := range l.levels {
			if k > currentLevel {
				continue
			}
			if int64(v.Threshold) <= totalAmount {
				continue
			}

			*refCoundToAdd -= v.Referrals
			lvls = append(lvls, k)
		}
	} else {
		for k, v := range l.levels {
			if k <= currentLevel {
				continue
			}
			if int64(v.Threshold) > totalAmount {
				continue
			}

			*refCoundToAdd += v.Referrals
			lvls = append(lvls, k)
		}
	}

	if len(lvls) == 0 {
		return refCoundToAdd, currentLevel
	}

	newLevel = slices.Max(lvls)
	if downgrade {
		newLevel = slices.Min(lvls) - 1
	}

	if l.levels[newLevel].Infinity {
		return nil, newLevel
	}
	return
}

func (l Levels) MinLvl() int {
	lvls := make([]int, 0, len(l.levels))
	for k := range l.levels {
		lvls = append(lvls, k)
	}

	// slices.Min will not panic because of previous logic
	return slices.Min(lvls)
}

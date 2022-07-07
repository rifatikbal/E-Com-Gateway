package config

import (
	"github.com/spf13/viper"
	"time"
)

var auth *AuthCfg

type AuthCfg struct {
	Secret   string
	HashCost uint64
	Duration time.Duration
}

func LoadAuthCfg() {
	auth = &AuthCfg{
		Secret:   viper.GetString("auth.secret"),
		HashCost: viper.GetUint64("auth.hash_cost"),
		Duration: viper.GetDuration("auth.duration") * time.Hour,
	}
}

func Auth() *AuthCfg {
	return auth
}

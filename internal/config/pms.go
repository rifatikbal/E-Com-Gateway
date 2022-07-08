package config

import (
	"github.com/spf13/viper"
)

var pms *PMSCfg

type PMSCfg struct {
	Url string
}

func LoadPMSCfg() {
	pms = &PMSCfg{
		Url: viper.GetString("PMS.url"),
	}
}

func PMS() *PMSCfg {
	return pms
}

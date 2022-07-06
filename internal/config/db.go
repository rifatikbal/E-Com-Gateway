package config

import (
	"time"

	"github.com/spf13/viper"
)

var db *DBCfg

type DBCfg struct {
	Host            string
	Port            int
	User            string
	Pass            string
	Name            string
	ConnMaxLifetime time.Duration
}

func LoadDBCfg() {
	db = &DBCfg{
		Host:            viper.GetString("database.host"),
		Port:            viper.GetInt("database.port"),
		User:            viper.GetString("database.user"),
		Pass:            viper.GetString("database.pass"),
		Name:            viper.GetString("database.name"),
		ConnMaxLifetime: viper.GetDuration("database.conn_max_lifetime") * time.Second,
	}
}

func DB() *DBCfg {
	return db
}

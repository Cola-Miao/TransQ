package config

import (
	"fmt"
	"github.com/Cola-Miao/TransQ/server/format"
	"github.com/spf13/viper"
	"log/slog"
)

var defaultCfg = config{
	Listener: listener{
		Timeout: 0,
	},
	LogLevel:    -4,
	ConnTimeout: 0,
}

func InitViper(cfgPath, cfgType string) error {
	format.FuncStart("InitViper")
	defer format.FuncEnd("InitViper")

	viper.AddConfigPath(cfgPath)
	viper.SetConfigType(cfgType)
	viper.SetDefault("server", defaultCfg)

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("viper.ReadInConfig: %v", err)
	}

	err = viper.UnmarshalKey("server", &Cfg)
	if err != nil {
		return fmt.Errorf("viper.UnmarshalKey: %w", err)
	}

	err = viper.SafeWriteConfig()
	if err != nil {
		slog.Warn("viper.SafeWriteConfig", "error", err.Error())
	}
	slog.Info("config", "server", Cfg)

	return nil
}

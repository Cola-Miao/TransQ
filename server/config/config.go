package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitViper(cfgPath, cfgType string) error {
	defaultCfg := config{
		Listener: listener{
			Timeout: 0,
		},
		LogLevel: -4,
	}

	viper.AddConfigPath(cfgPath)
	viper.SetConfigType(cfgType)
	viper.SetDefault("server", defaultCfg)

	err := viper.SafeWriteConfig()
	if err != nil {
		return fmt.Errorf("SafeWriteConfig: %w", err)
	}

	return nil
}

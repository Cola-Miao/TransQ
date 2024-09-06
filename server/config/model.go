package config

import "time"

var (
	Cfg config
)

type config struct {
	Listener listener `yaml:"listener"`
	Cache    cache    `yaml:"cache"`
	LogLevel int      `yaml:"logLevel"`
}

type listener struct {
	Timeout time.Duration `yaml:"timeout"`
}

type cache struct {
	DefaultExpiration time.Duration `yaml:"defaultExpiration"`
	CleanupInterval   time.Duration `yaml:"cleanupInterval"`
}

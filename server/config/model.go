package config

import "time"

var (
	Cfg config
)

type config struct {
	Listener listener `yaml:"listener"`
	LogLevel int      `yaml:"logLevel"`
}

type listener struct {
	Timeout time.Duration `yaml:"timeout"`
}

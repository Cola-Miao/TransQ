package config

var (
	Cfg config
)

type config struct {
	Listener listener `yaml:"listener"`
	LogLevel int      `yaml:"logLevel"`
}

type listener struct {
	Timeout int `yaml:"timeout"`
}

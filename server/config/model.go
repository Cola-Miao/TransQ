package config

var (
	Cfg config
)

type config struct {
	Listener    listener `yaml:"listener"`
	LogLevel    int      `yaml:"logLevel"`
	ConnTimeout int      `yaml:"connTimeout"` // Second, 0 never timeout
}

type listener struct {
	Timeout int `yaml:"timeout"`
}

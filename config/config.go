package config

// Config ...
type Config struct {
	URL     string
	Key     string
	Version string
}

// Newonfig ...
func NewConfig() *Config {
	return &Config{}
}

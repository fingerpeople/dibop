package config

// Config ...
type Config struct {
	URL string
	Key string
}

// Newonfig ...
func Newonfig() *Config {
	return &Config{}
}

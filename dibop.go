package dibop

import (
	"github.com/fingerpeople/dibop/config"
	"github.com/fingerpeople/dibop/sdk/user"
)

// Dibop ...
type Dibop struct {
	Urls    string
	Key     string
	Version string
}

// InputConfig ...
func InputConfig() *Dibop {
	return &Dibop{}
}

// Client ...
func Client(cfgData *Dibop) *config.Config {
	cfg := &config.Config{}
	cfg.Key = cfgData.Key
	cfg.URL = cfgData.Urls
	cfg.Version = cfgData.Version
	return cfg
}

// SDK ...
type SDK struct {
	Users user.UserInterface
}

// SDKHandler ...
func SDKHandler(cfg *config.Config) *SDK {
	return &SDK{
		Users: user.UserHandler(cfg),
	}
}

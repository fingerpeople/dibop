package dibop

import (
	"github.com/fingerpeople/dibop/config"
)

// Dibop ...
type Dibop struct {
	Urls string
	Key  string
}

// Client ...
func Client(config *config.Config) *Dibop {
	cfg := &Dibop{}
	cfg.Key = config.Key
	cfg.Urls = config.URL
	return cfg
}

// User ...
func (cl *Dibop) User() {

}

// Paramaters ...
func (cl *Dibop) Paramaters() {

}

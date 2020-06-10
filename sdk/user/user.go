package user

import (
	"fmt"

	"github.com/fingerpeople/dibop/config"
	"github.com/fingerpeople/dibop/utils/requester"
)

type User struct {
	Urls      string
	Key       string
	Version   string
	Requester requester.RequesterInterface
}

// UserHandler ....
func UserHandler(configs *config.Config) *User {
	return &User{
		Urls:      configs.URL,
		Key:       configs.Key,
		Version:   configs.Version,
		Requester: requester.RequesterHandler(),
	}
}

// UserInterface ...
type UserInterface interface {
	GetUserDetails()
}

// initURL ...
func (handler *User) initURL() string {
	urls := fmt.Sprintf("%s/%s/config/user", handler.Urls, handler.Version)
	return urls
}

// GetUserList ...
func (handler *User) GetUserList() {

}

// GetUserDetails ...
func (handler *User) GetUserDetails() {
	uri := handler.initURL()
	fmt.Println(uri)
}

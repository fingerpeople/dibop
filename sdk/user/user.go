package user

import (
	"encoding/json"
	"fmt"

	"github.com/fingerpeople/dibop/config"
	"github.com/fingerpeople/dibop/entity"
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
	GetUserDetails(user string) (*entity.ResponseUserDetails, error)
}

// initURL ...
func (handler *User) initURL() string {
	urls := fmt.Sprintf("%s/%s/config/user/", handler.Urls, handler.Version)
	return urls
}

func (handler *User) initHeader() map[string]string {
	return map[string]string{
		"Authorization": handler.Key,
	}
}

// GetUserList ...
func (handler *User) GetUserList() {

}

// GetUserDetails ...
func (handler *User) GetUserDetails(user string) (*entity.ResponseUserDetails, error) {
	uri := handler.initURL() + "get?user=" + user
	data, err := handler.Requester.GET(uri, handler.initHeader())
	if err != nil {
		return nil, err
	}
	result := &entity.ResponseUserDetails{}
	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

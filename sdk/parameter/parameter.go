package parameter

import (
	"encoding/json"
	"fmt"

	"github.com/fingerpeople/dibop/config"
	"github.com/fingerpeople/dibop/entity"
	"github.com/fingerpeople/dibop/utils/requester"
)

// Parameter ...
type Parameter struct {
	Urls      string
	Key       string
	Version   string
	Requester requester.RequesterInterface
}

// ParameterHandler ...
func ParameterHandler(configs *config.Config) *Parameter {
	return &Parameter{
		Urls:      configs.URL,
		Key:       configs.Key,
		Version:   configs.Version,
		Requester: requester.RequesterHandler(),
	}
}

// ParameterInterface ...
type ParameterInterface interface {
	PutRaw(params *entity.RawInput) (*entity.RawOutput, error)
}

// initURL ...
func (handler *Parameter) initURL() string {
	urls := fmt.Sprintf("%s/%s/", handler.Urls, handler.Version)
	return urls
}

func (handler *Parameter) initHeader() map[string]string {
	return map[string]string{
		"Authorization": handler.Key,
	}
}

// PutRaw ...
func (handler *Parameter) PutRaw(params *entity.RawInput) (*entity.RawOutput, error) {
	uri := handler.initURL() + "put/raw"
	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	data, err := handler.Requester.POST(uri, handler.initHeader(), payload)
	fmt.Println(string(data))
	return nil, nil
}

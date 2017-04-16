package emailacid

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"path"

	"github.com/parnurzeal/gorequest"
	"github.com/pkg/errors"
)

type EmailAcidClient struct {
	APIKey      string
	Password    string
	url         string
	ClientTypes []EmailClientType
}

type EmailAcidError struct {
	Error EmailAcidErrorBody `json:error:omitempty"`
}

type EmailAcidErrorBody struct {
	Name    string `json:"name:omitempty"`
	Message string `json:"message:omitempty"`
}

func New(APIKey, password string, clientTypes []EmailClientType) *EmailAcidClient {
	return &EmailAcidClient{
		APIKey:      APIKey,
		Password:    password,
		url:         "https://api.emailonacid.com/v5",
		ClientTypes: clientTypes,
	}
}

func (client *EmailAcidClient) buildRequest(method, resourcePath string) (*gorequest.SuperAgent, error) {
	request := gorequest.New()
	u, err := url.Parse(client.url)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, resourcePath)
	url := u.String()
	switch method {
	case gorequest.POST:
		request = request.Post(url)
	case gorequest.PUT:
		request = request.Put(url)
	case gorequest.DELETE:
		request = request.Delete(url)
	case gorequest.GET:
		request = request.Get(url)
	default:
		return nil, fmt.Errorf("invalid method %s", method)
	}
	request = request.SetBasicAuth(client.APIKey, client.Password)
	return request, nil
}

func sendRequest(request *gorequest.SuperAgent, in, out interface{}) error {
	res, body, errs := request.Send(in).EndStruct(out)
	if len(errs) != 0 {
		log.Printf("errors making request: %s", errs)
		return errs[0]
	}
	log.Printf("[%d] %s - %s", res.StatusCode, body, out)
	if res.StatusCode < 200 || res.StatusCode > 299 {
		var apiError EmailAcidError
		json.Unmarshal([]byte(body), &apiError)
		log.Printf("error: %s", apiError)
		return errors.Errorf("error making request: %s", apiError.Error.Message)
	}
	return nil
}

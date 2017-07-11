package emailacid

import (
	"encoding/json"
	"fmt"
	"net/url"
	"path"

	"github.com/parnurzeal/gorequest"
	"github.com/pkg/errors"
	"github.com/pressly/lg"
	"github.com/sirupsen/logrus"
)

type EmailAcidClient struct {
	APIKey      string
	Password    string
	url         string
	ClientTypes []ClientType
}

type EmailAcidError struct {
	Error EmailAcidErrorBody `json:error:omitempty"`
}

type EmailAcidErrorBody struct {
	Name    string `json:"name:omitempty"`
	Message string `json:"message:omitempty"`
}

func New(APIKey, password string, clientTypes []ClientType, verbose bool) *EmailAcidClient {
	logger := logrus.New()
	if verbose {
		logger.Level = logrus.DebugLevel
	}
	lg.DefaultLogger = logger
	lg.RedirectStdlogOutput(logger)
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

func sendRequest(request *gorequest.SuperAgent, in, out interface{}) (string, error) {
	lg.Debugf("req [%s] %s", request.Method, request.Url)
	res, body, errs := request.Send(in).EndStruct(out)
	if len(errs) != 0 {
		return "", errs[0]
	}
	lg.Debugf("res [%d]", res.StatusCode)
	if res.StatusCode < 200 || res.StatusCode > 299 {
		var apiError EmailAcidError
		json.Unmarshal([]byte(body), &apiError)
		return string(body), errors.Errorf("error making request: %s", apiError.Error.Message)
	}
	return string(body), nil
}

package emailacid

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/parnurzeal/gorequest"
)

type EmailTest struct {
	ID               string       `json:"id,omitempty"`
	Subject          string       `json:"subject,omitempty"`
	HTML             string       `json:"html,omitempty"`
	TransferEncoding string       `json:"transfer_encoding,omitempty"`
	Charset          string       `json:"charset,omitempty"`
	ReferenceID      string       `json:"reference_id,omitempty"`
	CustomerID       string       `json:"customer_id,omitempty"`
	Clients          []ClientType `json:"clients,omitempty"`
	ImageBlocking    bool         `json:"image_blocking,omitempty"`
}

type EmailTestList struct {
	Items []EmailTest
}

type EmailTestInfo struct {
	Subject    string   `json:"subject,omitempty"`
	Date       int64    `json:"date,omitempty"`
	Completed  []string `json:"completed,omitempty"`
	Processing []string `json:"processing,omitempty"`
	Bounced    []string `json:"bounced,omitempty"`
}

type EmailTestResult struct {
	ID            string                       `json:"id,omitempty"`
	DisplayName   string                       `json:"display_name,omitempty"`
	Client        string                       `json:"client,omitempty"`
	OS            string                       `json:"os,omitempty"`
	Category      string                       `json:"category,omitempty"`
	Screenshots   EmailTestResultScreenshots   `json:"screenshots,omitempty"`
	Thumbnail     string                       `json:"thumbnail,omitempty"`
	Status        string                       `json:"status,omitempty"`
	StatusDetails EmailTestResultStatusDetails `json:"status_details,omitempty"`
}

type EmailTestResultScreenshots struct {
	Default            string `json:"default,omitempty"`
	NoImages           string `json:"no_images,omitempty"`
	Horizontal         string `json:"horizontal,omitempty"`
	HorizontalNoImages string `json:"horizontal_no_images,omitempty"`
}

type EmailTestResultStatusDetails struct {
	Submitted     int64  `json:"submitted,omitempty"`
	Completed     int64  `json:"completed,omitempty"`
	BounceCode    int64  `json:"bounce_code,omitempty"`
	BounceMEssage string `json:"bounce_message,omitempty"`
}

type EmailTestResultList struct {
	Items EmailTestResultMap `json:"results,omitempty"`
}

type EmailTestResultMap map[string]EmailTestResult

func (client *EmailAcidClient) CreateTest(in *EmailTest) (*EmailTest, error) {
	request, err := client.buildRequest(gorequest.POST, "email/tests")
	if err != nil {
		return nil, err
	}
	var out EmailTest
	_, err = sendRequest(request, in, &out)
	return &out, err
}

func (client *EmailAcidClient) ListTests() (*EmailTestList, error) {
	request, err := client.buildRequest(gorequest.GET, "email/tests")
	if err != nil {
		return nil, err
	}
	var out EmailTestList
	_, err = sendRequest(request, nil, &out)
	return &out, err
}

func (client *EmailAcidClient) GetTestInfo(ID string) (*EmailTestInfo, error) {
	request, err := client.buildRequest(gorequest.GET, fmt.Sprintf("email/tests/%s", ID))
	if err != nil {
		return nil, err
	}
	var out EmailTestInfo
	_, err = sendRequest(request, nil, &out)
	return &out, err
}

func (client *EmailAcidClient) DeleteTest(ID string) (bool, error) {
	request, err := client.buildRequest(gorequest.DELETE, fmt.Sprintf("email/tests/%s", ID))
	if err != nil {
		return false, err
	}
	var out map[string]bool
	_, err = sendRequest(request, nil, &out)
	return out["success"], err
}

func (client *EmailAcidClient) GetTestClientResult(ID, clientID string) (*EmailTestResult, error) {
	request, err := client.buildRequest(gorequest.GET, fmt.Sprintf("email/tests/%s/results/%s", ID, clientID))
	if err != nil {
		return nil, err
	}
	var out EmailTestResultMap
	_, err = sendRequest(request, nil, &out)
	keys := reflect.ValueOf(out).MapKeys()
	emailTestResult := out[keys[0].String()]
	return &emailTestResult, err
}

func (client *EmailAcidClient) ListTestClientResults(ID string) (*EmailTestResultList, error) {
	request, err := client.buildRequest(gorequest.GET, fmt.Sprintf("email/tests/%s/results", ID))
	if err != nil {
		return nil, err
	}
	var out EmailTestResultList
	body, err := sendRequest(request, nil, nil)
	// Need to wrap the response into a json object before unmarshalling since the response is a object with objects
	body = fmt.Sprintf("{\"results\":%s}", body)
	err = json.Unmarshal([]byte(body), &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

/* TODO endpoints
Reprocess
Get test content
Client analysis
Link validation
Spam results
Spam seed list
*/

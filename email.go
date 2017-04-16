package emailacid

import "github.com/parnurzeal/gorequest"

type EmailTestReq struct {
	Subject          string            `json:"subject,omitempty"`
	HTML             string            `json:"html,omitempty"`
	TransferEncoding string            `json:"transfer_encoding,omitempty"`
	Charset          string            `json:"charset,omitempty"`
	ReferenceID      string            `json:"reference_id,omitempty"`
	CustomerID       string            `json:"customer_id,omitempty"`
	Clients          []EmailClientType `json:"clients,omitempty"`
	ImageBlocking    bool              `json:"image_blocking,omitempty"`
}

type EmailTestResList struct {
	EmailTests []EmailTestRes
}

type EmailTestRes struct {
	ID          string `json:"id,omitempty"`
	ReferenceID string `json:"reference_id,omitempty"`
	CustomerID  string `json:"customer_id,omitempty"`
}

func (client *EmailAcidClient) CreateTest(in *EmailTestReq) (*EmailTestRes, error) {
	request, err := client.buildRequest(gorequest.POST, "email/tests")
	if err != nil {
		return nil, err
	}
	var out EmailTestRes
	err = sendRequest(request, in, &out)
	return &out, err
}

func (client *EmailAcidClient) ListTests() (*EmailTestResList, error) {
	request, err := client.buildRequest(gorequest.GET, "email/tests")
	if err != nil {
		return nil, err
	}
	var out EmailTestResList
	err = sendRequest(request, nil, &out)
	return &out, err

}



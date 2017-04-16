package emailacid

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

type EmailClientType uint

const (
	Outlook2003 EmailClientType = iota
	Outlook2010
	GmailChrome22Win
)

var emailClientTypes = []string{
	"outlook03",
	"outlook10",
	"gmail_chr22_win",
}

type EmailClientMap map[string]EmailClient

type EmailClientResList struct {
	Clients EmailClientMap `json:"clients,omitempty"`
}

type EmailClient struct {
	ID            string `json:"id,omitempty"`
	Client        string `json:"client,omitempty"`
	OS            string `json:"os,omitempty"`
	Category      string `json:"category,omitempty"`
	ImageBlocking bool   `json:"image_blocking,omitempty"`
	Rotate        bool   `json:"rotate,omitempty"`
	Default       bool   `json:"default,omitempty"`
}

func (client *EmailAcidClient) ListClients() (*EmailClientResList, error) {
	request, err := client.buildRequest(gorequest.GET, "email/clients")
	if err != nil {
		return nil, err
	}
	var out EmailClientResList
	err = sendRequest(request, nil, &out)
	return &out, err
}

func (t EmailClientType) String() string {
	return emailClientTypes[t]
}

func (t EmailClientType) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t *EmailClientType) UnmarshalText(text []byte) error {
	enum := string(text)
	for i := 0; i < len(emailClientTypes); i++ {
		if enum == emailClientTypes[i] {
			*t = EmailClientType(i)
			return nil
		}
	}
	return fmt.Errorf("unknown email template type: %s", enum)
}

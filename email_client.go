package emailacid

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

type ClientType uint

const (
	Android5_1 ClientType = iota
	Android6
	Android7
	AppleMail8
	AppleMail9
	AppleMail10
	GmailAppIOS
	IPadAir9
	IPad3_9
	IPadAir10
	IPadMini10
	IPadPro10
	IPadRetina10
	IPadMini9
	IPhoneSE10
	IPhone5_9
	IPhone6_9
	IPhone6p_9
	IPhone6_10
	IPhone6p_10
	IPhone7_10
	IPhone7p_10
	Notes6_5
	Notes7
	Notes8
	Notes8_5
	Outlook03
	Outlook07
	Outlook10
	Outlook10D125
	OutlookMac11
	Outlook13
	Outlook13D125
	Outlook16
	OutlookMac16
	Thunderbird13
	AolChrome26Win
	AolChrome26Mac
	AolFirefox21Win
	AolFirefox21Mac
	AolIE10Win
	AolIE11Win
	AolSafari8Mac
	BolChrome26Win
	ComcastChrome26Win
	ComcastFirefox21Win
	FFRChrome26Win
	FreenetChrome26Win
	GmailChrome26Win
	GmailChrome26Mac
	GmailFirefox21Mac
	GmailFirefox21Win
	GmailIE10Win
	GmailIE11Win
	GmailSafari8Mac
	GMXChrome26Win
	GAppsChrome26Win
	GAppsFirefox21Win
	GAppsIO10Win
	GAppsIE11Win
	LaposteChrome26Win
	LiberoChrome26Win
	MailRuChrome26Win
	NateChrome26Win
	NaverChrome26Win
	Office365Chrome26Win
	Office365Firefox21Win
	Office365IE10Win
	Office365IE11Win
	OFRChrome26Win
	OutlookChrome26Win
	OutlookChrome26Mac
	OutlookFirefox21Mac
	OutlookFirefox21Win
	OutlookIE10Win
	OutlookIE11Win
	OutlookSafari8Mac
	SFRChrome26Win
	TOEChrome26Win
	TelestraChrome26Win
	TerraChrome26Win
	WDEChrome26Win
	YBChrome26Win
	YBChrome26Mac
	YBFirefox21Win
	TBFirefox21Mac
	IBIE10Win
	YBIE11Win
	YDXChrome26Win
)

var ClientTypes = []string{
	"android5_1",
	"android6",
	"android7",
	"applemail8",
	"applemail9",
	"applemail10",
	"gmailapp_ios",
	"ipadair_9",
	"ipad3_9",
	"ipadair_10",
	"ipadmini_10",
	"ipadpro_10",
	"ipadretina_10",
	"ipadmini_9",
	"iphonese_10",
	"iphone5_9",
	"iphone6_9",
	"iphone6p_9",
	"iphone6_10",
	"iphone6p_10",
	"iphone7_10",
	"iphone7p_10",
	"notes65",
	"notes7",
	"notes8",
	"notes85",
	"outlook03",
	"outlook07",
	"outlook10",
	"outlook10_d125",
	"macoutlook11",
	"outlook13",
	"outlook13_d125",
	"outlook16",
	"macoutlook16",
	"thunderbird13",
	"aol_chr26_win",
	"aol_chr26_mac",
	"aol_ff21_win",
	"aol_ff21_mac",
	"aol_ie10_win",
	"aol_ie11_win",
	"aol_si8_mac",
	"bol_chr26_win",
	"cc_chr26_win",
	"cc_ff21_win",
	"ffr_chr26_win",
	"frt_chr26_win",
	"gmail_chr26_win",
	"gmail_chr26_mac",
	"gmail_ff21_mac",
	"gmail_ff21_win",
	"gmail_ie10_win",
	"gmail_ie11_win",
	"gmail_si8_mac",
	"gmx_chr26_win",
	"gapps_chr26_win",
	"gapps_ff21_win",
	"gapps_ie10_win",
	"gapps_ie11_win",
	"lpe_chr26_win",
	"lb_chr26_win",
	"mru_chr26_win",
	"nate_chr26_Win",
	"naver_chr26_Win",
	"o365_chr26_win",
	"o365_ff21_win",
	"o365_ie10_win",
	"o365_ie11_win",
	"ofr_chr26_win",
	"ol_chr26_win",
	"ol_chr26_mac",
	"ol_ff21_mac",
	"ol_ff21_win",
	"ol_ie10_win",
	"ol_ie11_win",
	"ol_si8_mac",
	"sfr_chr26_win",
	"toe_chr26_win",
	"telstra_chr26_Win",
	"tra_chr26_win",
	"wde_chr26_win",
	"yb_chr26_win",
	"yb_chr26_mac",
	"yb_ff21_win",
	"yb_ff21_mac",
	"yb_ie10_win",
	"yb_ie11_win",
	"ydx_chr26_Win",
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

type EmailClientTypes struct {
	IDs []ClientType `json:"clients,omitempty"`
}

func (client *EmailAcidClient) ListClients() (*EmailClientResList, error) {
	request, err := client.buildRequest(gorequest.GET, "email/clients")
	if err != nil {
		return nil, err
	}
	var out EmailClientResList
	_, err = sendRequest(request, nil, &out)
	return &out, err
}

func (client *EmailAcidClient) ListDefaultClientIDs() (*EmailClientTypes, error) {
	request, err := client.buildRequest(gorequest.GET, "email/clients/default")
	if err != nil {
		return nil, err
	}
	var out EmailClientTypes
	_, err = sendRequest(request, nil, &out)
	return &out, err
}

func (t ClientType) String() string {
	return ClientTypes[t]
}

func (t ClientType) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t *ClientType) UnmarshalText(text []byte) error {
	enum := string(text)
	for i := 0; i < len(ClientTypes); i++ {
		if enum == ClientTypes[i] {
			*t = ClientType(i)
			return nil
		}
	}
	return fmt.Errorf("unknown email template type: %s", enum)
}

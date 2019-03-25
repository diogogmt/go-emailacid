Golang Email on Acid API client
------------------------
go-emailacid is a [Go](http://golang.org/) client package for accessing the [Email on Acid](https://www.emailonacid.com) [API](https://api.emailonacid.com/docs/latest/overview).

<a href="http://golang.org"><img alt="Go package" src="https://golang.org/doc/gopher/appenginegophercolor.jpg" width="20%" /></a>
<a href="http://trello.com"><img src="https://www.emailonacid.com/images/logos/white-splat.png" style="height: 80px; margin-bottom: 2em;"></a>

[![GoDoc](https://godoc.org/github.com/pressly/go-emailacid?status.png)](https://godoc.org/github.com/pressly/go-emailacid)


## Examples

Create a new client

```golang
clientTypes := []emailacid.ClientType{
  emailacid.Outlook03,
  emailacid.Outlook10,
}
EmailAcidClient = emailacid.New(apiKey, password, clientTypes)
```

Submit a new test

```golang
in := &emailacid.EmailTest{
    Subject:          "test",
    HTML:             "<html><body>testing...</body></html>",
    ReferenceID:      "12345ABC",
    CustomerID:       "2",
    Clients:          []emailacid.ClientType{
      emailacid.Outlook03,
      emailacid.Outlook10,
    }
  }
_, err := EmailAcidClient.CreateTest(in)
```

Get test results

```golang
testID := "sandbox"
clientID := "outlook03"
_, err := EmailAcidClient.GetTestClientResult(testID, clientID)
```

The test result contains the following information:
```golang
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
```


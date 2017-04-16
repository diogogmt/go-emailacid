package emailacid_test

import (
	"testing"

	emailacid "github.com/diogogmt/go-emailacid"
)

func TestEmail(t *testing.T) {
	t.Parallel()
	t.Run("createTest", testEmailCreateTest)
	t.Run("listTests", testEmailCreateTest)
}

func testEmailCreateTest(t *testing.T) {
	in := &emailacid.EmailTestReq{
		Subject:          "test",
		HTML:             "<html><body>testing...</body></html>",
		TransferEncoding: "8bit",
		Charset:          "utf-8",
		ReferenceID:      "12345ABC",
		CustomerID:       "2",
		Clients:          EmailAcidClient.ClientTypes,
	}
	_, err := EmailAcidClient.CreateTest(in)
	if err != nil {
		t.Fatal(err)
	}
}

func testEmailListTests(t *testing.T) {
	_, err := EmailAcidClient.ListTests()
	if err != nil {
		t.Fatal(err)
	}
}

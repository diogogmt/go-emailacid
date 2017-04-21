package emailacid_test

import (
	"testing"

	emailacid "github.com/diogogmt/go-emailacid"
)

func TestEmail(t *testing.T) {
	t.Parallel()
	t.Run("createTest", testEmailCreateTest)
	t.Run("listTests", testEmailCreateTest)
	t.Run("getTestInfo", testEmailGetTestInfo)
	t.Run("getTestClientResult", testGetTestClientResult)
	t.Run("listTestClientResults", testListTestClientResults)
}

func testEmailCreateTest(t *testing.T) {
	in := &emailacid.EmailTest{
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

func testEmailGetTestInfo(t *testing.T) {
	testID := "sandbox"
	_, err := EmailAcidClient.GetTestInfo(testID)
	if err != nil {
		t.Fatal(err)
	}
}

func testGetTestClientResult(t *testing.T) {
	testID := "sandbox"
	clientID := "outlook03"
	_, err := EmailAcidClient.GetTestClientResult(testID, clientID)
	if err != nil {
		t.Fatal(err)
	}
}

func testListTestClientResults(t *testing.T) {
	testID := "sandbox"
	_, err := EmailAcidClient.ListTestClientResults(testID)
	if err != nil {
		t.Fatal(err)
	}
}

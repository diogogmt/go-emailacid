package emailacid_test

import (
	"testing"
)

func TestEmailClient(t *testing.T) {
	t.Parallel()
	t.Run("listClients", testEmailClientList)
}

func testEmailClientList(t *testing.T) {
	_, err := EmailAcidClient.ListClients()
	if err != nil {
		t.Fatal(err)
	}
}

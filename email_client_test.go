package emailacid_test

import (
	"testing"
)

func TestEmailClient(t *testing.T) {
	t.Parallel()
	t.Run("listClients", testListClients)
	t.Run("listDefaultClientIDs", testListDefaultClientIDs)
}

func testListClients(t *testing.T) {
	_, err := EmailAcidClient.ListClients()
	if err != nil {
		t.Fatal(err)
	}
}

func testListDefaultClientIDs(t *testing.T) {
	_, err := EmailAcidClient.ListDefaultClientIDs()
	if err != nil {
		t.Fatal(err)
	}
}

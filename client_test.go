package emailacid_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/pressly/go-emailacid"
)

var EmailAcidClient *emailacid.EmailAcidClient

func TestMain(m *testing.M) {
	start := time.Now()

	apiKey := os.Getenv("EMAILACID_API_KEY")
	if apiKey == "" {
		apiKey = "sandbox"
	}
	password := os.Getenv("EMAILACID_PASSWORD")
	if password == "" {
		password = "sandbox"
	}
	clientTypes := []emailacid.ClientType{
		emailacid.Outlook03,
		emailacid.Outlook10,
		emailacid.GmailChrome26Win,
	}
	EmailAcidClient = emailacid.New(apiKey, password, clientTypes)

	exitCode := m.Run()

	log.Printf("=== EXIT \t\t\t\t%v\n", time.Since(start))
	os.Exit(exitCode)
}

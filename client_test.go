package emailacid_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/diogogmt/go-emailacid"
)

var EmailAcidClient *emailacid.EmailAcidClient

func TestMain(m *testing.M) {
	start := time.Now()

	apiKey := os.Getenv("EMAILACID_API_KEY")
	password := os.Getenv("EMAILACID_PASSWORD")
	clientTypes := []emailacid.EmailClientType{
		emailacid.Outlook2003,
		emailacid.Outlook2010,
		emailacid.GmailChrome22Win,
	}
	EmailAcidClient = emailacid.New(apiKey, password, clientTypes)

	exitCode := m.Run()

	log.Printf("=== EXIT \t\t\t\t%v\n", time.Since(start))
	os.Exit(exitCode)
}

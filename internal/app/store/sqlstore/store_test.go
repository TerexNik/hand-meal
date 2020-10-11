package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=192.168.20.160 dbname=hand_meal sslmode=disable user=postgres password=postgres"
	}

	os.Exit(m.Run())
}

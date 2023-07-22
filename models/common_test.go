package models

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	ConnectTestDatabase()
	os.Exit(m.Run())
}

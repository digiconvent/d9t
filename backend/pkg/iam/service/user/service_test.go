package iam_user_service_test

import (
	"testing"

	"github.com/digiconvent/migrate_packages/db"
)

var testDB db.DatabaseInterface

func TestMain(m *testing.M) {
	m.Run()
}

package package_databases_test

import (
	"os"
	"path"
	"testing"

	"github.com/digiconvent/d9t/meta/environment"
	"github.com/digiconvent/d9t/meta/package_databases"
)

func TestGetMigratedDatabases(t *testing.T) {
	dataPath := path.Join(os.TempDir(), "package_databases_test")
	os.RemoveAll(dataPath)

	environment.Env.InstalledVersion = "-1.-1.-1"
	migrateFrom, err := package_databases.MigrateDatabasesFrom("../../pkg")
	if err != nil {
		t.Fatal("Did not expect err:" + err.Error())
	}
	_, err = migrateFrom.To(dataPath)
	if err != nil {
		t.Fatal("Did not expect an error, instead got ", err.Error())
	}
}

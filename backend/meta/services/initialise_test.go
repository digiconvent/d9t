package services_test

import (
	"os"
	"path"
	"testing"

	"github.com/digiconvent/d9t/meta/environment"
	"github.com/digiconvent/d9t/meta/package_databases"
	"github.com/digiconvent/d9t/meta/services"
)

func TestInitialiseServices(t *testing.T) {
	dataPath := path.Join(os.TempDir(), "test_services")
	os.RemoveAll(dataPath)
	environment.Env.InstalledVersion = "-1.-1.-1" // prerequisite
	migrations, err := package_databases.MigrateDatabasesFrom("../../pkg")
	if err != nil {
		t.Fatal("need databases, instead got err", err)
	}

	databases, err := migrations.To(dataPath)
	if err != nil {
		t.Fatalf("did not expect err, instead got %v", err.Error())
	}

	services, err := services.Initialise(databases)
	if err != nil {
		t.Fatal("did not expect err, instead got", err)
	}
	if services.Iam == nil {
		t.Fatal("did not expect Iam service to be nil")
	}
}

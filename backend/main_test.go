package main_test

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/digiconvent/migrate_packages"
)

func TestMain(m *testing.M) {
	// remove test databases from previous tests
	thisFolder, _ := os.Getwd()
	testFolder := path.Join(os.TempDir(), "digiconvent_test")
	os.RemoveAll(testFolder)

	// create the test databases
	mgr, err := migrate_packages.FromSemVer(-1, -1, -1).ToVersion(-1, -1, -1).Verbose().WithLocalFilesAt(thisFolder, "pkg")
	if err != nil {
		fmt.Println(err)
	}
	_, err = mgr.MigrateDatabasesIn(testFolder)
	if err != nil {
		fmt.Println(err)
	}

}

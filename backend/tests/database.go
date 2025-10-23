package tests

import (
	"fmt"
	"os"
	"path"
	"sync"

	"github.com/digiconvent/migrate_packages"
	"github.com/digiconvent/migrate_packages/db"
)

var databases map[string]db.DatabaseInterface

var once sync.Once
var mu sync.Mutex

func initializeDatabases() {
	fmt.Println("This is executed only once!")
	// Use PID to create unique folder per test binary, preventing conflicts
	testDataFolder := path.Join(os.TempDir(), fmt.Sprintf("digiconvent_test_%d", os.Getpid()))
	os.RemoveAll(testDataFolder)

	projectRoot := findProjectRoot()

	mgr, err := migrate_packages.FromSemVer(-1, -1, -1).ToVersion(-1, -1, -1).WithLocalFilesAt(projectRoot, "pkg")
	if err != nil {
		panic(err)
	}

	// MigrateDatabasesIn will either create new databases or connect to existing ones
	databases, err = mgr.MigrateDatabasesIn(testDataFolder)
	if err != nil {
		panic(err)
	}
}

func GetTestDatabase(packageName string) db.DatabaseInterface {
	once.Do(initializeDatabases)

	database, exists := databases[packageName]
	if !exists {
		panic("Test database not found for package: " + packageName)
	}
	return database
}

func findProjectRoot() string {
	currentDir, _ := os.Getwd()

	for {
		if _, err := os.Stat(path.Join(currentDir, "go.mod")); err == nil {
			return currentDir
		}

		parent := path.Dir(currentDir)
		if parent == currentDir {

			panic("Could not find project root (go.mod not found)")
		}
		currentDir = parent
	}
}

package package_databases

import (
	"fmt"
	"os"

	"github.com/digiconvent/d9t/meta/environment"
	"github.com/digiconvent/d9t/utils/ffs"
	"github.com/digiconvent/migrate_packages"
	"github.com/digiconvent/migrate_packages/db"
)

type chooseDestinationFolder interface {
	To(dataPath string) (map[string]db.DatabaseInterface, error)
}

func MigrateDatabasesFrom(pkgPath string) (chooseDestinationFolder, error) {
	if !ffs.Exists(pkgPath) {
		return nil, fmt.Errorf("pkg path %v does not exist", pkgPath)
	}
	return &migratePkgs{
		pkgDir: pkgPath,
	}, nil
}

type migratePkgs struct {
	pkgDir  string
	dataDir string
}

func (m *migratePkgs) To(dataPath string) (map[string]db.DatabaseInterface, error) {
	m.dataDir = dataPath
	return m.migrate()
}

func (m *migratePkgs) migrate() (map[string]db.DatabaseInterface, error) {
	installedVersion := migrate_packages.ToVersion(environment.Env.InstalledVersion)
	targetVersion := migrate_packages.ToVersion(environment.BinaryVersion)
	migration := migrate_packages.From(installedVersion).To(targetVersion).Verbose()

	var migrationManager migrate_packages.PackageManager
	if targetVersion.Major == -1 {
		projectRoot, _ := os.Getwd()
		migrationManager, _ = migration.WithLocalFilesAt(projectRoot, m.pkgDir)
	} else {
		repoManager, err := migration.WithPublicRepository("digiconvent", "d9t")
		if err != nil {
			return nil, err
		}
		migrationManager, err = repoManager.WithPkgDir(m.pkgDir)
		if err != nil {
			return nil, err
		}
	}

	databases, err := migrationManager.MigrateDatabasesIn(m.dataDir)
	if err != nil {
		return nil, err
	}

	if installedVersion.IsUninitialised() || !installedVersion.Equals(targetVersion) {
		// safe to say that at this point all the migrations went through so we can set the version to targetVersion
		fmt.Printf("Succesfully migrated from %v to %v\n", environment.Env.InstalledVersion, environment.BinaryVersion)
		environment.Env.InstalledVersion = targetVersion.String()
		environment.Save()
	}
	return databases, nil
}

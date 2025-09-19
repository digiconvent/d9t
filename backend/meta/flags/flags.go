package flags

import (
	"errors"
	"flag"
	"os"
	"slices"
	"strings"

	"github.com/digiconvent/d9t/meta/environment"
	"github.com/digiconvent/d9t/utils/logging"
	"github.com/digiconvent/embed_env"
	"github.com/digiconvent/migrate_packages"
)

var allowedModes = []string{"install", "serve", "env", "update"}

var flagSet = flag.NewFlagSet("flags", flag.ContinueOnError)
var Mode = flagSet.String("mode", "", "Mode to run this binary ["+strings.Join(allowedModes, "/")+"]\n"+installationExplanation())
var Force = flagSet.Bool("force", false, "Use to override")
var EnvPreset = flagSet.String("preset", "", "Environment preset (this is advanced, don't use this if you're not familiar)")
var Test bool = false

func ProcessFlags(args []string) error {
	// set Test flag in case tests are running
	for _, flag := range args {
		if strings.HasPrefix(flag, "-test.testlogfile") {
			Test = true
			break
		}
	}
	err := flagSet.Parse(args)
	if err != nil {
		return err
	}

	if !slices.Contains(allowedModes, *Mode) {
		if !Test {
			flagSet.Usage()
		}
		return errors.New("'" + *Mode + "' is not a valid option for --mode")
	}

	return nil
}

func installationExplanation() string {
	x := "install\t\t"
	installedBinary := "/home/digiconvent/main"
	installedEnv := &environment.EnvVars{}
	thisVersion := migrate_packages.ToVersion(environment.BinaryVersion)
	if _, err := os.Stat(installedBinary); err == nil {
		err := embed_env.ReadFromBinary(installedBinary, installedEnv, "")
		installedVersion := migrate_packages.ToVersion(installedEnv.InstalledVersion)
		if err != nil {
			panic(err)
		}
		if thisVersion.EarlierThan(installedVersion) {
			x += "Since " + installedVersion.String() + " is installed and this binary has version " + environment.BinaryVersion + ", you cannot use this option"
			x = logging.Disabled(x)
		} else if thisVersion.Equals(installedVersion) {
			x += "This version is already installed"
			x = logging.Disabled(x)
		} else {
			x += "Use this to upgrade from " + installedEnv.InstalledVersion + " to " + environment.BinaryVersion
			x = logging.Green(x)
		}
	} else {
		x += "Use this to install a fresh digiconvent:" + environment.BinaryVersion
	}
	return x
}

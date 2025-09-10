package flags

import (
	"errors"
	"flag"
	"slices"
	"strings"
)

var allowedModes = []string{"install"}

var flagSet = flag.NewFlagSet("flags", flag.ContinueOnError)
var Mode = flagSet.String("mode", "", "mode to run this binary ["+strings.Join(allowedModes, "/")+"]")
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

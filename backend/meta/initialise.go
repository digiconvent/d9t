package meta

import (
	"github.com/digiconvent/d9t/meta/environment"
	"github.com/digiconvent/d9t/meta/flags"
)

func Initialise(call []string) error {
	err := flags.ProcessFlags(call[1:])
	if err != nil {
		return err
	}

	err = environment.Load(*flags.EnvPreset)
	if err != nil {
		return err
	}

	return nil
}

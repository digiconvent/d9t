package flags_test

import (
	"testing"

	"github.com/digiconvent/d9t/meta/flags"
)

func TestProcessFlags(t *testing.T) {
	t.Run("mode", func(t *testing.T) {
		var allowedFlags = []string{"install"}

		for _, allowedFlag := range allowedFlags {
			err := flags.ProcessFlags([]string{"--mode", allowedFlag})
			if err != nil {
				t.Fatal("expected --mode", allowedFlag, "to be allowed")
			}
		}
	})
}

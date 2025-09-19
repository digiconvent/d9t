package ffs_test

import (
	"os"
	"path"
	"testing"

	"github.com/digiconvent/d9t/utils/ffs"
)

func TestExists(t *testing.T) {
	if !ffs.Exists(os.TempDir()) {
		t.Fatalf("expected %v to exist", os.TempDir())
	}
	x := path.Join(os.TempDir(), "randomfilename")
	if ffs.Exists(x) {
		t.Fatalf("did not expect %v to exist", x)
	}
}

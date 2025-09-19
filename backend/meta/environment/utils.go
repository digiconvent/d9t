package environment

import (
	"os"
	"path"
)

func Pwd() string {
	return path.Dir(Uri())
}

func Uri() string {
	u, _ := os.Executable()
	return u
}

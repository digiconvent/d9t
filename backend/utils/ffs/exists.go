package ffs

import "os"

func Exists(uri string) bool {
	_, err := os.Stat(uri)
	return err == nil
}

package core_test

import (
	"testing"

	"github.com/digiconvent/d9t/core"
)

func TestStatus(t *testing.T) {
	var isError map[int]bool = map[int]bool{
		200: false,
		201: false,
		204: false,
		400: true,
		401: true,
		403: true,
		404: true,
		409: true,
		422: true,
		500: true,
	}

	for val := range isError {
		status := core.Status{
			Code: val,
		}
		if status.Err() != isError[val] {
			t.Fatal("expected", val, " to be err")
		}
	}
}

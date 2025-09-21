package context

import (
	"encoding/json"
	"net/http"

	"github.com/digiconvent/d9t/core"
	"github.com/go-playground/validator/v10"
)

func ParseAndValidate[T any](r *http.Request) (*T, *core.Status) {
	var payload *T

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return payload, core.UnprocessableContentError("Could not convert body to json")
	}
	defer r.Body.Close()

	v := &validator.Validate{}
	if err := v.Struct(payload); err != nil {
		return payload, core.UnprocessableContentError(err.Error())
	}

	return payload, nil
}

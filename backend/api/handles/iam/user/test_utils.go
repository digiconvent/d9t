package iam_user_handles

import (
	"bytes"
	"net/http/httptest"
	"sync"

	"github.com/digiconvent/d9t/api/context"
	"github.com/digiconvent/d9t/meta/services"
	"github.com/digiconvent/d9t/tests"
	"github.com/digiconvent/migrate_packages/db"
	"github.com/google/uuid"
)

var (
	testServices *services.Services
	once         sync.Once
)

func getTestServices() *services.Services {
	once.Do(func() {
		databases := map[string]db.DatabaseInterface{"iam": tests.GetTestDatabase("iam")}
		var err error
		testServices, err = services.Initialise(databases)
		if err != nil {
			panic("Failed to initialize test services: " + err.Error())
		}
	})
	return testServices
}

func newTestContext(method, body string) *context.Context {
	req := httptest.NewRequest(method, "/", bytes.NewReader([]byte(body)))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	id := uuid.New()

	return &context.Context{
		Id:       &id,
		Request:  req,
		Response: resp,
		Services: getTestServices(),
	}
}
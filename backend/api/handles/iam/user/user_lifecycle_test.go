package iam_user_handles

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func TestUserLifecycle(t *testing.T) {
	// 1. create user
	createCtx := newTestContext("POST", `{"email":"lifecycle@example.com"}`)
	Create(createCtx)

	createRec := createCtx.Response.(*httptest.ResponseRecorder)
	if createRec.Code != 201 {
		t.Fatalf("Create failed: expected 201, got %d: %s", createRec.Code, createRec.Body.String())
	}

	userIdStr := createRec.Body.String()
	if !strings.Contains(userIdStr, "-") {
		t.Fatal("Create should return UUID")
	}

	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		t.Fatalf("Invalid UUID returned: %s", userIdStr)
	}

	// 2. read user to verify creation
	readCtx := newTestContext("GET", "")
	readCtx.Id = &userId
	Read(readCtx)

	readRec := readCtx.Response.(*httptest.ResponseRecorder)
	if readRec.Code != 200 {
		t.Fatalf("expected 200, instead got %v", readRec.Code)
	}

	jsonRes := readRec.Body.Bytes()
	parsedUser := iam_domain.User{}
	err = json.Unmarshal(jsonRes, &parsedUser)
	if err != nil {
		t.Fatalf("cannot unmarshal %v: %v", string(jsonRes), err)
	}

	// 3. update user
	randomFirstName, _ := uuid.NewV7()
	randomLastName, _ := uuid.NewV7()
	updateCtx := newTestContext("PUT", `{"first_name": "`+randomFirstName.String()+`", "last_name": "`+randomLastName.String()+`"}`)
	updateCtx.Id = &userId
	Update(updateCtx)

	updateRec := updateCtx.Response.(*httptest.ResponseRecorder)
	if updateRec.Code != 204 {
		t.Error("expected 204 instead got " + updateRec.Result().Status)
	}

	// 4. list users to verify user exists
	listCtx := newTestContext("GET", "")
	List(listCtx)

	listRec := listCtx.Response.(*httptest.ResponseRecorder)
	jsonRes = listRec.Body.Bytes()
	parsed := []iam_domain.UserProxy{}
	json.Unmarshal(jsonRes, &parsed)
	found := false
	if len(parsed) != 2 {
		t.Error("Expected 2 users")
	} else {
		for _, entry := range parsed {
			if entry.Id.String() == userId.String() {
				found = true
				if entry.FirstName != randomFirstName.String() || entry.LastName != randomLastName.String() {
					t.Fatal("expected names to be updated")
				}
			}
		}
	}

	if !found {
		t.Fatalf("expected user %v to be found", userId)
	}

	// 5. delete user
	deleteCtx := newTestContext("DELETE", `{}`)
	deleteCtx.Id = &userId
	Delete(deleteCtx)

	deleteRec := deleteCtx.Response.(*httptest.ResponseRecorder)
	if deleteRec.Code != 204 {
		t.Fatalf("expected 204, instead got %v", deleteRec.Code)
	}

	// 6. verify user is deleted by attempting to read
	readAfterDeleteCtx := newTestContext("GET", "")
	readAfterDeleteCtx.Id = &userId
	Read(readAfterDeleteCtx)

	readAfterDeleteRec := readAfterDeleteCtx.Response.(*httptest.ResponseRecorder)
	if readAfterDeleteRec.Code != 404 {
		t.Fatalf("expected 404 user not found but instead got %v, %v", readAfterDeleteRec.Code, readAfterDeleteRec.Body)
	}
}

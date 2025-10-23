package iam_user_handles

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateValidation(t *testing.T) {
	tests := []struct {
		name           string
		body           string
		expectedStatus int
		expectedError  string
	}{
		{
			name:           "missing email",
			body:           `{}`,
			expectedStatus: 422,
			expectedError:  "validation",
		},
		{
			name:           "invalid json",
			body:           `{invalid}`,
			expectedStatus: 422,
			expectedError:  "json",
		},
		{
			name:           "empty email",
			body:           `{"email":""}`,
			expectedStatus: 422,
			expectedError:  "validation",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := newTestContext("POST", tt.body)
			Create(ctx)

			rec := ctx.Response.(*httptest.ResponseRecorder)
			if rec.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d: %s", tt.expectedStatus, rec.Code, rec.Body.String())
			}

			if tt.expectedError != "" && !strings.Contains(rec.Body.String(), tt.expectedError) {
				t.Errorf("expected error containing '%s', got: %s", tt.expectedError, rec.Body.String())
			}
		})
	}
}

func TestCreateDuplicateEmail(t *testing.T) {
	email := "duplicate@e2e.com"

	ctx1 := newTestContext("POST", `{"email":"`+email+`"}`)
	Create(ctx1)

	rec1 := ctx1.Response.(*httptest.ResponseRecorder)
	if rec1.Code != 201 {
		t.Fatalf("First create failed: %d - %s %s", rec1.Code, rec1.Body.String(), email)
	}

	ctx2 := newTestContext("POST", `{"email":"`+email+`"}`)
	Create(ctx2)

	rec2 := ctx2.Response.(*httptest.ResponseRecorder)
	if rec2.Code != 409 {
		t.Errorf("Expected conflict (409) for duplicate email, got %d: %s", rec2.Code, rec2.Body.String())
	}

	if !strings.Contains(rec2.Body.String(), "iam.user.email.duplicate") {
		t.Error("Should indicate email conflict")
	}
}

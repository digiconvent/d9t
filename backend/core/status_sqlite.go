package core

import (
	"database/sql"
	"net/http"
)

func ErrToCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case sql.ErrNoRows:
		return http.StatusNotFound
	}

	if isConstraintError(err) {
		return http.StatusConflict
	}

	if isTimeoutError(err) {
		return http.StatusRequestTimeout
	}

	if isConnectionError(err) {
		return http.StatusServiceUnavailable
	}

	if isSyntaxError(err) {
		return http.StatusBadRequest
	}

	return http.StatusInternalServerError
}

func isConstraintError(err error) bool {
	return containsAny(err.Error(), []string{
		"UNIQUE constraint failed",
		"FOREIGN KEY constraint failed",
		"NOT NULL constraint failed",
		"constraint failed",
		"constraint violation",
	})
}

func isTimeoutError(err error) bool {
	return containsAny(err.Error(), []string{
		"timeout",
		"busy",
		"database is locked",
	})
}

func isConnectionError(err error) bool {
	return containsAny(err.Error(), []string{
		"unable to open database file",
		"no such database",
		"disk I/O error",
		"database disk image is malformed",
	})
}

func isSyntaxError(err error) bool {
	return containsAny(err.Error(), []string{
		"syntax error",
		"near \"",
		"no such table",
		"no such column",
	})
}

func containsAny(s string, substrs []string) bool {
	for _, substr := range substrs {
		if contains(s, substr) {
			return true
		}
	}
	return false
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[:len(substr)] == substr
}

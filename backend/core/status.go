package core

import (
	"fmt"

	"github.com/digiconvent/d9t/utils/logging"
)

type Status struct {
	Message string
	Code    int
}

func BadRequestError(message string) *Status {
	logging.Error("Bad Request Error: " + message)
	return &Status{
		Message: message,
		Code:    400,
	}
}

func UnauthorizedError(message string) *Status {
	logging.Error("Unauthorized: " + message)
	return &Status{
		Message: message,
		Code:    401,
	}
}

func ForbiddenError(message string) *Status {
	logging.Error("Forbidden: " + message)
	return &Status{
		Message: message,
		Code:    403,
	}
}

func NotFoundError(message string) *Status {
	logging.Error("Not Found Error: " + message)
	return &Status{
		Message: message,
		Code:    404,
	}
}

func ConflictError(message string) *Status {
	logging.Error("Conflict: " + message)
	return &Status{
		Message: message,
		Code:    409,
	}
}

func UnprocessableContentError(message string) *Status {
	logging.Error("Could not process content: " + message)
	return &Status{
		Message: message,
		Code:    422,
	}
}

func InternalError(message string) *Status {
	logging.Error("Internal Error: " + message)
	return &Status{
		Message: message,
		Code:    500,
	}
}

func IsProcessing() *Status {
	return &Status{
		Message: "Processing",
		Code:    102,
	}
}

func StatusSuccess() *Status {
	return &Status{
		Message: "Success",
		Code:    200,
	}
}

func StatusCreated() *Status {
	return &Status{
		Message: "Created",
		Code:    201,
	}
}

func StatusNoContent() *Status {
	return &Status{
		Message: "No Content",
		Code:    204,
	}
}

func (e *Status) Ok() bool {
	if e == nil {
		return false
	}
	return e.Code == 200 || e.Code == 201 || e.Code == 204
}

func (e *Status) Err() bool {
	return !e.Ok()
}

func (e *Status) String() string {
	return fmt.Sprint(e.Code, ": ", e.Message)
}

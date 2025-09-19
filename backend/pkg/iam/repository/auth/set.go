package iam_auth_repository

import (
	"github.com/digiconvent/d9t/core"
	uuid "github.com/google/uuid"
)

func (r *IamAuthRepository) SetCredentialEmailaddress(id *uuid.UUID, email string) *core.Status {
	panic("unimplemented")
}

func (r *IamAuthRepository) SetCredentialPassword(id *uuid.UUID, password string) *core.Status {
	if id == nil {
		return core.UnprocessableContentError("ID is required")
	}
	result, err := r.db.Exec("update users set password = ? where id = ?", password, id.String())
	if err != nil {
		return core.InternalError(err.Error())
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return core.NotFoundError("User not found")
	}
	return core.StatusNoContent()
}

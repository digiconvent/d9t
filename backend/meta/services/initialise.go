package services

import (
	"errors"

	iam_repository "github.com/digiconvent/d9t/pkg/iam/repo"
	iam_service "github.com/digiconvent/d9t/pkg/iam/service"
	"github.com/digiconvent/migrate_packages/db"
)

type Services struct {
	Iam *iam_service.IamServices
}

var Ref *Services

func Initialise(databases map[string]db.DatabaseInterface) (*Services, error) {
	if Ref == nil {
		iamDatabase := databases["iam"]
		if iamDatabase == nil {
			return nil, errors.New("could not find iam database")
		}
		Ref = &Services{
			Iam: iam_service.NewIamServices(iam_repository.NewIamRepository(iamDatabase)),
		}
	}

	return Ref, nil
}

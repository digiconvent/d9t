package iam_service_test

import (
	"os"
	"path"
	"testing"

	"github.com/digiconvent/d9t/meta/environment"
	iam_repository "github.com/digiconvent/d9t/pkg/iam/repository"
	iam_service "github.com/digiconvent/d9t/pkg/iam/service"
	"github.com/digiconvent/migrate_packages/db"
	"github.com/google/uuid"
)

func NewTestIamService() *iam_service.IamServices {
	testDb, _ := db.New(path.Join(os.TempDir(), "digiconvent_test", "iam", "iam.db"))
	iamRepo := iam_repository.NewIamRepository(testDb)
	return iam_service.NewIamServices(iamRepo)
}

func getRootPermissionGroup() string {
	testService := GetTestIamService()
	facades, _ := testService.PermissionGroup.ListPermissionGroups()
	for _, facade := range facades {
		if facade.Name == "root" {
			return facade.Id.String()
		}
	}

	return "" // id.String()
}

func GetRootPermissionGroupUuid() *uuid.UUID {
	root := getRootPermissionGroup()
	res := uuid.MustParse(root)
	return &res
}

func GetTestIamService() *iam_service.IamServices {
	env := environment.FromString("jwt_private_key=-----BEGIN+RSA+PRIVATE+KEY-----%0AMIICXQIBAAKBgQC63Z3rSU2BJGjqBJ2m2foORciYSM4WxCmZC0sM3WYPsWvVSJFL%0AioLhmYI41bRVjsR5iLODjIHoJGfS7MhwIuBuNSrPH%2Fw0WY6N9wsYHJP3ysv%2FwE7p%0AMWP6yTdtvMQUrvej1tdJgeTH9WUptZWyL9wFznNuKAmvohdB%2FagubL55tQIDAQAB%0AAoGAOC6ZZDxNyR4hTPrGS3wN6myguuhozf98Yj86%2BUs3fFBjk2SWCphCcYGt3Vij%0AN3YNLF0lW8%2FK3vjTHl%2BCg%2FR4dSaRP9pfMe5a%2BKq6Q5JZ45qtQO4tab7pINFpxmHO%0A4fCskIP1%2FUHRIOlytiqvJCs6ba95%2B4CswGyJFZeGMX%2BpBKECQQDcUUt%2FDR1K7na0%0ABoRG4AqdzSqrkC%2FhA6Wvo%2FQRr8bQGGTWM3CVlVeChNlWrJHBiY%2BT%2Fr%2F9anYAPVOe%0ApVAXAXlfAkEA2SFbJkyUIHBlWW4zbIzfSuskIiyxC%2FKTvd5q3MwfT99Vp3r72U%2BB%0A%2BxedNwiH7DWot7Sq6O3%2B%2FotuGnd5VbShawJBANQmB%2F%2BwnitgUrdpqwgghriLLnKD%0A5kbTX0ExLD23uEb7pmXIpLm1U83fpPTcOYJWtwX4geFBGZ2DcDjM0MFGbxsCQQCV%0AEcOEfdrMkp3iUjhX9DdxNhKuq8Q5kadBGayAPlvY%2BFnUDClQPrfcbY%2FwMpku4IAe%0A75jPJmaE17EV1qSk5MhVAkBdpxL5M4Pkkk020ynGztsFFdEWe7XdQywNip8hJOoq%0AexD%2BxB7Q0hULBD7B9MZPeZyDn2j26M6oTCB5DTLo0EzX%0A-----END+RSA+PRIVATE+KEY-----%0A")
	environment.Env = env
	testDb, _ := db.New(path.Join(os.TempDir(), "digiconvent_test", "iam", "iam.db"))
	repo := iam_repository.NewIamRepository(testDb)
	s := iam_service.NewIamServices(repo)
	return s
}

func TestMain(m *testing.M) {
	GetTestIamService()
	m.Run()
}

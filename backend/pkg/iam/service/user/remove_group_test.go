package iam_user_service_test

import (
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	iam_repository "github.com/digiconvent/d9t/pkg/iam/repo"
	iam_group_service "github.com/digiconvent/d9t/pkg/iam/service/group"
	iam_user_service "github.com/digiconvent/d9t/pkg/iam/service/user"
	"github.com/digiconvent/d9t/tests"
)

func TestRemoveGroup(t *testing.T) {
	db := tests.GetTestDatabase("iam")
	repo := iam_repository.NewIamRepository(db)
	userService := iam_user_service.NewUserService(repo.User)
	groupService := iam_group_service.NewGroupService(repo.Group)

	user := &iam_domain.User{
		Email:     "removegroup@example.com",
		FirstName: "RemoveGroup",
		LastName:  "Test",
		Enabled:   true,
	}
	userId, _ := userService.Create(user)

	group := &iam_domain.Group{
		Name:   "Test Group",
		Type:   "role",
		Parent: getGroupRoot(),
	}
	groupId, _ := groupService.Create(group)

	userService.AddGroup(userId, groupId)

	status := userService.RemoveGroup(userId, groupId)
	if !status.Ok() {
		t.Fatalf("RemoveGroup failed: %s", status.String())
	}
}

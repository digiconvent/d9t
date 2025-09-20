package iam_group_service_test

import (
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	iam_repository "github.com/digiconvent/d9t/pkg/iam/repo"
	iam_group_service "github.com/digiconvent/d9t/pkg/iam/service/group"
	iam_user_service "github.com/digiconvent/d9t/pkg/iam/service/user"
	"github.com/digiconvent/d9t/tests"
)

func TestAddUser(t *testing.T) {
	db := tests.GetTestDatabase("iam")
	repo := iam_repository.NewIamRepository(db)
	groupService := iam_group_service.NewGroupService(repo.Group)
	userService := iam_user_service.NewUserService(repo.User)

	group := &iam_domain.Group{Name: "Add User Group", Type: "role", Parent: getGroupRoot()}
	groupId, _ := groupService.Create(group)

	user := &iam_domain.User{Email: "adduser@example.com", FirstName: "Add", LastName: "User", Enabled: true}
	userId, _ := userService.Create(user)

	status := groupService.AddUser(groupId, userId)
	if !status.Ok() {
		t.Fatalf("AddUser failed: %s", status.String())
	}
}

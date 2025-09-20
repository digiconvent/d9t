package iam_group_service_test

import (
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/digiconvent/d9t/tests"
)

func TestRemoveUser(t *testing.T) {
	groupService := tests.TestGroupService()
	userService := tests.TestUserService()

	group := &iam_domain.Group{Name: "Remove User Group", Type: "role", Parent: getGroupRoot()}
	groupId, _ := groupService.Create(group)

	user := &iam_domain.User{Email: "removeuser@example.com", FirstName: "Remove", LastName: "User", Enabled: true}
	userId, _ := userService.Create(user)

	groupService.AddUser(groupId, userId)

	status := groupService.RemoveUser(groupId, userId)
	if !status.Ok() {
		t.Fatalf("RemoveUser failed: %s", status.String())
	}
}

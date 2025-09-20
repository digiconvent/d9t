package iam_group_service_test

import (
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	iam_repository "github.com/digiconvent/d9t/pkg/iam/repo"
	iam_group_service "github.com/digiconvent/d9t/pkg/iam/service/group"
	iam_user_service "github.com/digiconvent/d9t/pkg/iam/service/user"
	"github.com/digiconvent/d9t/tests"
)

func TestGroupCyclePreventionSelfParent(t *testing.T) {
	service := tests.TestGroupService()

	group := &iam_domain.Group{
		Name:   "Self Parent Test",
		Type:   "container",
		Parent: getGroupRoot(),
	}

	groupId, _ := service.Create(group)

	status := service.SetParent(groupId, groupId)
	if status.Ok() {
		t.Fatal("Expected self-parent assignment to fail")
	}
}

func TestGroupCyclePreventionDirectCycle(t *testing.T) {
	service := tests.TestGroupService()

	groupA := &iam_domain.Group{Name: "Group A", Type: "container", Parent: getGroupRoot()}
	groupAId, _ := service.Create(groupA)

	groupB := &iam_domain.Group{Name: "Group B", Type: "container", Parent: getGroupRoot()}
	groupBId, _ := service.Create(groupB)

	status := service.SetParent(groupBId, groupAId)
	if !status.Ok() {
		t.Fatalf("First parent assignment failed: %s", status.String())
	}

	status = service.SetParent(groupAId, groupBId)
	if status.Ok() {
		t.Fatal("Expected cycle creation to fail")
	}
}

func TestGroupCyclePreventionMultiLevel(t *testing.T) {
	service := tests.TestGroupService()

	groupA := &iam_domain.Group{Name: "Group A", Type: "container", Parent: getGroupRoot()}
	groupAId, _ := service.Create(groupA)

	groupB := &iam_domain.Group{Name: "Group B", Type: "container", Parent: getGroupRoot()}
	groupBId, _ := service.Create(groupB)

	groupC := &iam_domain.Group{Name: "Group C", Type: "container", Parent: getGroupRoot()}
	groupCId, _ := service.Create(groupC)

	status := service.SetParent(groupBId, groupAId)
	if !status.Ok() {
		t.Fatalf("A→B parent assignment failed: %s", status.String())
	}

	status = service.SetParent(groupCId, groupBId)
	if !status.Ok() {
		t.Fatalf("B→C parent assignment failed: %s", status.String())
	}

	status = service.SetParent(groupAId, groupCId)
	if status.Ok() {
		t.Fatal("Expected multi-level cycle creation to fail")
	}
}

func TestContainerGroupCannotHaveUsers(t *testing.T) {
	db := tests.GetTestDatabase("iam")
	repo := iam_repository.NewIamRepository(db)
	groupService := iam_group_service.NewGroupService(repo.Group)
	userService := iam_user_service.NewUserService(repo.User)

	containerGroup := &iam_domain.Group{
		Name:   "Container Group",
		Type:   "container",
		Parent: getGroupRoot(),
	}
	groupId, _ := groupService.Create(containerGroup)

	user := &iam_domain.User{
		Email:     "container@example.com",
		FirstName: "Container",
		LastName:  "Test",
		Enabled:   true,
	}
	userId, _ := userService.Create(user)

	status := groupService.AddUser(groupId, userId)
	if status.Ok() {
		t.Fatal("Expected adding user to container group to fail")
	}
}

func TestRoleGroupCanHaveUsers(t *testing.T) {
	db := tests.GetTestDatabase("iam")
	repo := iam_repository.NewIamRepository(db)
	groupService := iam_group_service.NewGroupService(repo.Group)
	userService := iam_user_service.NewUserService(repo.User)

	roleGroup := &iam_domain.Group{
		Name:   "Role Group",
		Type:   "role",
		Parent: getGroupRoot(),
	}
	groupId, _ := groupService.Create(roleGroup)

	user := &iam_domain.User{
		Email:     "role@example.com",
		FirstName: "Role",
		LastName:  "Test",
		Enabled:   true,
	}
	userId, _ := userService.Create(user)

	status := groupService.AddUser(groupId, userId)
	if !status.Ok() {
		t.Fatalf("Failed to add user to role group: %s", status.String())
	}
}

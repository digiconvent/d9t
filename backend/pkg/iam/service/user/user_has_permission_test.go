package iam_user_service_test

import (
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/digiconvent/d9t/tests"
	"github.com/google/uuid"
)

func TestUserHasPermission_ComplexScenario(t *testing.T) {
	userService := tests.TestUserService()
	groupService := tests.TestGroupService()
	policyService := tests.TestPolicyService()

	policyAId, _ := policyService.Create(&iam_domain.Policy{
		Name:        "PolicyA",
		Description: "Policy A",
	})

	policyBId, _ := policyService.Create(&iam_domain.Policy{
		Name:        "PolicyB",
		Description: "Policy B",
	})

	policyCId, _ := policyService.Create(&iam_domain.Policy{
		Name:        "PolicyC",
		Description: "Policy C",
	})

	policyService.AddPermission(policyAId, "iam.users.read")
	policyService.AddPermission(policyBId, "iam.users.read")
	policyService.AddPermission(policyBId, "iam.users.create")
	policyService.AddPermission(policyCId, "iam.groups.read")

	groupAId, _ := groupService.Create(&iam_domain.Group{
		Name:   "GroupA",
		Type:   "role",
		Parent: getGroupRoot(),
	})

	groupBId, _ := groupService.Create(&iam_domain.Group{
		Name:   "GroupB",
		Type:   "role",
		Parent: getGroupRoot(),
	})

	groupCId, _ := groupService.Create(&iam_domain.Group{
		Name:   "GroupC",
		Type:   "role",
		Parent: getGroupRoot(),
	})

	groupService.AddPolicy(groupAId, policyAId)
	groupService.AddPolicy(groupBId, policyBId)
	groupService.AddPolicy(groupCId, policyAId)
	groupService.AddPolicy(groupCId, policyCId)

	userAId, _ := userService.Create(&iam_domain.User{
		Email:     "usera@test.com",
		FirstName: "User",
		LastName:  "A",
		Enabled:   true,
	})

	userBId, _ := userService.Create(&iam_domain.User{
		Email:     "userb@test.com",
		FirstName: "User",
		LastName:  "B",
		Enabled:   true,
	})

	userCId, _ := userService.Create(&iam_domain.User{
		Email:     "userc@test.com",
		FirstName: "User",
		LastName:  "C",
		Enabled:   true,
	})
	userService.AddGroup(userAId, groupAId)
	userService.AddGroup(userBId, groupBId)
	userService.AddGroup(userCId, groupAId)
	userService.AddGroup(userCId, groupCId)

	t.Run("UserA_SingleGroup", func(t *testing.T) {
		policies, status := userService.UserHasPermission(userAId, "iam.users.read")
		if status.Err() {
			t.Fatalf("Failed to check users.read: %s", status.String())
		}
		if len(policies) != 1 || policies[0].Name != "PolicyA" {
			t.Errorf("UserA should have iam.users.read through PolicyA, got %d policies", len(policies))
		}

		policies, status = userService.UserHasPermission(userAId, "iam.users.create")
		if status.Err() {
			t.Fatalf("Failed to check users.create: %s", status.String())
		}
		if len(policies) != 0 {
			t.Errorf("UserA should not have iam.users.create, got %d policies", len(policies))
		}
	})

	t.Run("UserB_DifferentPermissions", func(t *testing.T) {
		policies, status := userService.UserHasPermission(userBId, "iam.users.read")
		if status.Err() {
			t.Fatalf("Failed to check users.read: %s", status.String())
		}
		if len(policies) != 1 || policies[0].Name != "PolicyB" {
			t.Errorf("UserB should have iam.users.read through PolicyB, got %d policies", len(policies))
		}

		policies, status = userService.UserHasPermission(userBId, "iam.users.create")
		if status.Err() {
			t.Fatalf("Failed to check users.create: %s", status.String())
		}
		if len(policies) != 1 || policies[0].Name != "PolicyB" {
			t.Errorf("UserB should have iam.users.create through PolicyB, got %d policies", len(policies))
		}
	})

	t.Run("UserC_MultipleGroups", func(t *testing.T) {
		policies, status := userService.UserHasPermission(userCId, "iam.users.read")
		if status.Err() {
			t.Fatalf("Failed to check users.read: %s", status.String())
		}
		if len(policies) != 1 || policies[0].Name != "PolicyA" {
			t.Errorf("UserC should have iam.users.read through PolicyA, got %d policies", len(policies))
		}

		policies, status = userService.UserHasPermission(userCId, "iam.groups.read")
		if status.Err() {
			t.Fatalf("Failed to check groups.read: %s", status.String())
		}
		if len(policies) != 1 || policies[0].Name != "PolicyC" {
			t.Errorf("UserC should have iam.groups.read through PolicyC, got %d policies", len(policies))
		}

		policies, status = userService.UserHasPermission(userCId, "iam.users.create")
		if status.Err() {
			t.Fatalf("Failed to check users.create: %s", status.String())
		}
		if len(policies) != 0 {
			t.Errorf("UserC should not have iam.users.create, got %d policies", len(policies))
		}
	})

	t.Run("EdgeCases", func(t *testing.T) {
		nonExistentUser := uuid.New()
		policies, status := userService.UserHasPermission(&nonExistentUser, "iam.users.read")
		if status.Err() {
			t.Fatalf("Failed to check permission for non-existent user: %s", status.String())
		}
		if len(policies) != 0 {
			t.Errorf("Non-existent user should have no policies, got %d", len(policies))
		}

		policies, status = userService.UserHasPermission(userAId, "iam.nonexistent.read")
		if status.Err() {
			t.Fatalf("Failed to check non-existent permission: %s", status.String())
		}
		if len(policies) != 0 {
			t.Errorf("Non-existent permission should return no policies, got %d", len(policies))
		}
	})
}

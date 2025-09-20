package iam_group_service_test

import (
	"github.com/digiconvent/d9t/tests"
	"github.com/google/uuid"
)

func getGroupRoot() *uuid.UUID {
	groups, _ := tests.TestGroupService().ReadProxies()
	for _, group := range groups {
		if group.Name == "root" {
			return group.Id
		}
	}
	return nil
}

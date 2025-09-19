package iam_user_repository

import (
	"slices"

	"github.com/digiconvent/d9t/core"
	pagination "github.com/digiconvent/d9t/core/page"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
)

var allowedSortBy = []string{"last_name", "first_name", "emailaddress"}

func (r *IamUserRepository) List(filter *iam_domain.UserFilterSort) (*pagination.Page[iam_domain.UserFacade], *core.Status) {
	sortClause := ""
	if filter != nil && slices.Contains(allowedSortBy, filter.Sort.Field) {
		sortClause = "order by " + filter.Sort.Field
		if !filter.Sort.Asc {
			sortClause += " desc"
		} else {
			sortClause += " asc"
		}
	}

	filterClause := ""

	users := []*iam_domain.UserFacade{}
	rows, err := r.db.Query("select id, first_name, last_name, status_id, status_name, role_id, role_name from user_facades " + sortClause)

	if err != nil {
		return nil, core.InternalError(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		user := iam_domain.UserFacade{}

		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.StatusId, &user.StatusName, &user.RoleId, &user.RoleName)

		if err != nil {
			return nil, core.InternalError(err.Error())
		}

		users = append(users, &user)
	}

	pageNumber := 1
	if filter != nil {
		pageNumber = filter.Page
	}

	itemsPerPage := 10
	if filter != nil {
		itemsPerPage = filter.ItemsPerPage
	}
	var page = &pagination.Page[iam_domain.UserFacade]{
		Items:        users,
		Page:         pageNumber,
		ItemsPerPage: itemsPerPage,
	}

	err = r.db.QueryRow("select count(*) from user_facades " + filterClause).Scan(&page.ItemsCount)

	if err != nil {
		return nil, core.InternalError(err.Error())
	}

	return page, core.StatusSuccess()
}

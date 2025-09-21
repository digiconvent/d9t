package api_engine

import iam_user_handles "github.com/digiconvent/d9t/api/handles/iam/user"

var routes = RouteTable{
	"iam": {
		"user": {
			"create": Post(iam_user_handles.Create),
			"read":   Get(iam_user_handles.Read),
			"update": Put(iam_user_handles.Update),
			"delete": Delete(iam_user_handles.Delete),
			"list":   Get(iam_user_handles.List),
		},
	},
}

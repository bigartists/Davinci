package view

import "davinci/module/user"

type Service interface {
	GetViews(projectId int64, u *user.User) (*[]ViewBaseInfo, error)
	GetViewDetail(u *user.User, id int64) (*ViewWithSourceBaseInfo, error)
	ExecuteSql(u *user.User, sourceId int64, sql string, variables []interface{}, Limit int64) (map[string]interface{}, error)
}

package view

import "davinci/module/user"

type Service interface {
	GetViews(projectId int64, u *user.User) (*[]ViewBaseInfo, error)
	GetViewDetail(u *user.User, id int64) (*ViewWithSourceBaseInfo, error)
}

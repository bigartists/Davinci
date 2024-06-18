package role

import "davinci/module/user"

type Service interface {
	GetRolesByProjectId(projectId int64, u *user.User) ([]RoleBaseInfo, error)
}

//RoleBaseInfo

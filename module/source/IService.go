package source

import "davinci/module/user"

type Service interface {
	GetSources(projectId int64, u *user.User) (*[]Source, error)
}

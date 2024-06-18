package project

import (
	"context"
	"davinci/module/user"
)

type Service interface {
	GetProjects(ctx context.Context, u *user.User) ([]ProjectWithCreateBy, error)
	GetFavoriteProjects(u *user.User) ([]ProjectWithCreateBy, error)
	GetProjectInfo(projectId int64, u *user.User) (ProjectInfo, error)
	GetProjectDetail(projectId int64, u *user.User, modify bool) (*ProjectDetail, error)
}

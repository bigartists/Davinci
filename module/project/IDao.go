package project

type Dao interface {
	GetProjectsByUser(id int64) ([]ProjectWithCreateBy, error)
	GetFavoriteProjects(id int64) ([]ProjectWithCreateBy, error)
	GetProjectDetail(id int64) (*ProjectDetail, error)
}

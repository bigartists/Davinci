package source

type Dao interface {
	GetSourcesByProject(projectId int64) (*[]Source, error)
}

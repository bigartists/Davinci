package relProjectAdmin

type Dao interface {
	GetByProjectAndUser(projectId int64, userId int64) (*RelProjectAdmin, error)
}

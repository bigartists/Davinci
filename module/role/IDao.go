package role

type Dao interface {
	getRoleBaseInfoByProject(projectId int64) ([]RoleBaseInfo, error)
}

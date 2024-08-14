package view

type Dao interface {
	getViewBaseInfoByProject(projectId int64) (*[]ViewBaseInfo, error)
	GetViewWithSourceBaseInfo(viewId int64) (*ViewWithSource, error)
}

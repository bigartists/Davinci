package relRoleView

type Dao interface {
	GetByView(viewId int64) (*[]RelRoleView, error)
}

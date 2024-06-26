package relRoleView

import "davinci/client"

var DaoGetter Dao

type DaoGetterImpl struct {
}

func (d DaoGetterImpl) GetByView(viewId int64) (*[]RelRoleView, error) {
	RelRoleView := &[]RelRoleView{}
	err := client.Orm.Where("view_id = ?", viewId).Find(&RelRoleView).Error
	if err != nil {
		return nil, err
	}
	return RelRoleView, nil
}

func init() {
	DaoGetter = NewDaoGetterImpl()
}

func NewDaoGetterImpl() *DaoGetterImpl {
	return &DaoGetterImpl{}
}

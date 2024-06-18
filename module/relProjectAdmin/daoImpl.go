package relProjectAdmin

import (
	"davinci/client"
	"fmt"
)

var DaoGetter Dao

type DaoGetterImpl struct {
}

func init() {
	DaoGetter = NewDaoGetterImpl()
}

func NewDaoGetterImpl() *DaoGetterImpl {
	return &DaoGetterImpl{}
}

func (i *DaoGetterImpl) GetByProjectAndUser(projectId int64, userId int64) (*RelProjectAdmin, error) {
	var relProjectAdmin RelProjectAdmin

	err := client.Orm.Table("rel_project_admin").
		Where("user_id = ? AND project_id = ?", userId, projectId).
		First(&relProjectAdmin).Error

	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return &relProjectAdmin, nil
}

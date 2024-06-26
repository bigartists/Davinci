package view

import (
	"davinci/client"
)

var DaoGetter Dao

type DaoGetterImpl struct {
}

func (d DaoGetterImpl) GetViewWithSourceBaseInfo(viewId int64) (*ViewWithSource, error) {
	viewWithSource := &ViewWithSource{}
	err := client.Orm.Table("view v").
		Select(
			"v.*, s.id as 'source.id', s.name as 'source.name'").
		Joins("left join source s on s.id=v.source_id").
		Where("v.id=?", viewId).
		Find(viewWithSource).Error
	if err != nil {
		return nil, err
	}
	return viewWithSource, nil
}

func (d DaoGetterImpl) getViewBaseInfoByProject(projectId int64) (*[]ViewBaseInfo, error) {

	viewBaseInfo := &[]ViewBaseInfo{}
	err := client.Orm.
		Table("view v").
		Select("v.id, v.`name`, v.`description`, s.name as 'sourceName'").
		Joins("left join source s on s.id = v.source_id").
		Where("v.project_id = ?", projectId).
		Find(viewBaseInfo).Error
	if err != nil {
		return nil, err
	}

	return viewBaseInfo, nil
}

func init() {
	DaoGetter = NewDaoGetterImpl()
}

func NewDaoGetterImpl() *DaoGetterImpl {
	return &DaoGetterImpl{}
}

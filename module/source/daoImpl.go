package source

import "davinci/client"

var DaoGetter Dao

type DaoGetterImpl struct {
}

func (d DaoGetterImpl) GetSourceById(id int64) (*Source, error) {
	var source Source
	err := client.Orm.Where("id = ?", id).Find(&source).Error
	return &source, err
}

func (d DaoGetterImpl) GetSourcesByProject(projectId int64) (*[]Source, error) {
	var sources []Source
	err := client.Orm.Where("project_id = ?", projectId).Find(&sources).Error
	return &sources, err
}

func init() {
	DaoGetter = NewDaoGetterImpl()
}

func NewDaoGetterImpl() *DaoGetterImpl {
	return &DaoGetterImpl{}
}

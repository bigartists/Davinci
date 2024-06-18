package role

import "davinci/client"

var DaoGetter Dao

type DaoGetterImpl struct {
}

func init() {
	DaoGetter = NewDaoGetterImpl()
}

func NewDaoGetterImpl() *DaoGetterImpl {
	return &DaoGetterImpl{}
}

func (d DaoGetterImpl) getRoleBaseInfoByProject(projectId int64) ([]RoleBaseInfo, error) {

	var roles []RoleBaseInfo
	err := client.Orm.Table("role r").
		Select("r.id, r.name, r.description").
		Joins("left join rel_role_project rrp on rrp.role_id = r.id").
		Where("rrp.project_id = ?", projectId).
		Scan(&roles).Error

	if err != nil {
		return nil, err
	}

	return roles, nil
}

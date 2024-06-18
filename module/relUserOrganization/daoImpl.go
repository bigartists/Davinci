package relUserOrganization

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

func (i *DaoGetterImpl) GetRel(userId int64, orgId int64) (*RelUserOrganization, error) {
	var relUserOrg RelUserOrganization

	err := client.Orm.Table("rel_user_organization").
		Where("user_id = ? AND org_id = ?", userId, orgId).
		First(&relUserOrg).Error

	if err != nil {
		return nil, err
	}
	return &relUserOrg, nil

}

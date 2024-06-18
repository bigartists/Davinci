package organization

import "davinci/client"

var DaoGetter Dao

type DaoGetterImpl struct {
}

func (d DaoGetterImpl) GetOrgMembers(orgId string) ([]*OrganizationWithUser, error) {
	db := client.Orm
	var members []*OrganizationWithUser
	err := db.Select("ruo.id, u.id AS 'user.id', IF(u.`name` is NULL,u.username,u.`name`) AS 'user.username', u.email AS 'user.email', u.avatar AS 'user.avatar', ruo.role AS 'user.role'").
		Table("user u").
		Joins("LEFT JOIN rel_user_organization ruo on ruo.user_id = u.id").
		Joins("LEFT JOIN organization o on o.id = ruo.org_id").
		Where("ruo.org_id = ?", orgId).
		Find(&members).Error

	if err != nil {
		return nil, err
	}
	return members, nil
}

func init() {
	DaoGetter = NewDaoGetterImpl()
}

func NewDaoGetterImpl() *DaoGetterImpl {
	return &DaoGetterImpl{}
}

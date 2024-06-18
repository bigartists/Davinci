package organization

import (
	"context"
	"davinci/client"
	"davinci/module/user"
)

var ServiceGetter Service

type ServiceGetterImpl struct{}

func (s ServiceGetterImpl) GetOrgMembers(orgId string) ([]*OrganizationWithUser, error) {
	members, err := DaoGetter.GetOrgMembers(orgId)
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (s ServiceGetterImpl) GetOrganizations(ctx context.Context, u *user.User) ([]*OrganizationInfo, error) {

	query := `
        SELECT o.*, IFNULL(ruo.role, 0) as role
        FROM organization o
        LEFT JOIN rel_user_organization ruo ON (ruo.org_id = o.id AND ruo.user_id = ?)
        WHERE o.id IN (
            SELECT id FROM organization WHERE user_id = ?
            UNION
            SELECT org_id as id FROM rel_user_organization WHERE user_id = ?
        )
    `
	var organizations []*OrganizationInfo
	if err := client.Orm.WithContext(ctx).Raw(query, u.Id, u.Id, u.Id).Find(&organizations).Error; err != nil {
		return nil, err
	}
	return organizations, nil
}

func NewServiceGetterImpl() *ServiceGetterImpl {
	return &ServiceGetterImpl{}
}

func init() {
	ServiceGetter = NewServiceGetterImpl()
}

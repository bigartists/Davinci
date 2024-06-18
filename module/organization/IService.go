package organization

import (
	"context"
	"davinci/module/user"
)

type Service interface {
	GetOrganizations(ctx context.Context, u *user.User) ([]*OrganizationInfo, error)
	GetOrgMembers(orgId string) ([]*OrganizationWithUser, error)
}

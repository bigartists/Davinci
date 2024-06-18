package organization

type Dao interface {
	GetOrgMembers(orgId string) ([]*OrganizationWithUser, error)
}

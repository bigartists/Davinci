package relUserOrganization

type Dao interface {
	GetRel(userId int64, orgId int64) (*RelUserOrganization, error)
}

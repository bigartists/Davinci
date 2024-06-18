package role

import (
	"davinci/module/user"
)

var ServiceGetter Service

func init() {
	ServiceGetter = NewServiceGetterImpl()
}

type ServiceGetterImpl struct{}

func NewServiceGetterImpl() *ServiceGetterImpl {
	return &ServiceGetterImpl{}
}

func (s ServiceGetterImpl) GetRolesByProjectId(projectId int64, u *user.User) ([]RoleBaseInfo, error) {
	// todo 校验projectid 是否存在

	list, err := DaoGetter.getRoleBaseInfoByProject(projectId)
	return list, err
}

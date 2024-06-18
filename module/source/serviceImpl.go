package source

import (
	"davinci/module/project"
	"davinci/module/user"
	"fmt"
)

var ServiceGetter Service

func init() {
	ServiceGetter = NewServiceGetterImpl()
}

type ServiceGetterImpl struct{}

func (s ServiceGetterImpl) GetSources(projectId int64, u *user.User) (*[]Source, error) {
	projectdetail, _ := project.ServiceGetter.GetProjectDetail(projectId, u, false)

	fmt.Print(projectdetail)

	sources, err := DaoGetter.GetSourcesByProject(projectId)

	return sources, err
}

func NewServiceGetterImpl() *ServiceGetterImpl {
	return &ServiceGetterImpl{}
}

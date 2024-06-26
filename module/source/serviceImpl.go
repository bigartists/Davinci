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

func (s ServiceGetterImpl) GetSourceDetail(u *user.User, id int64) (*SourceDetail, error) {
	source, err := s.getSource(id)
	if err != nil {
		return nil, err
	}
	// TODO ： 校验当前用户是否有查看权限 缺失了，后续补充上

	sourceDto := ConvertSourcesToSourceDetail(source)

	return sourceDto, nil
}

func (s ServiceGetterImpl) getSource(id int64) (*Source, error) {
	source, err := DaoGetter.GetSourceById(id)
	return source, err
}

func (s ServiceGetterImpl) GetSources(projectId int64, u *user.User) (*[]Source, error) {
	projectDetail, _ := project.ServiceGetter.GetProjectDetail(projectId, u, false)

	fmt.Print(projectDetail)

	sources, err := DaoGetter.GetSourcesByProject(projectId)

	return sources, err
}

func NewServiceGetterImpl() *ServiceGetterImpl {
	return &ServiceGetterImpl{}
}

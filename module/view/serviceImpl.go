package view

import (
	"davinci/module/project"
	"davinci/module/relRoleView"
	"davinci/module/source"
	"davinci/module/user"
	"fmt"
)

var ServiceGetter Service

func init() {
	ServiceGetter = NewServiceGetterImpl()
}

type ServiceGetterImpl struct{}

func (s ServiceGetterImpl) ExecuteSql(u *user.User, sourceId int64, sql string, variables []interface{}, Limit int64) (map[string]interface{}, error) {
	rows, err := source.ServiceGetter.ExecuteSql(u, sourceId, sql, variables, Limit)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (s ServiceGetterImpl) GetViewDetail(u *user.User, id int64) (*ViewWithSourceBaseInfo, error) {

	view, err := DaoGetter.GetViewWithSourceBaseInfo(id)
	fmt.Println("view=,", view)

	if err != nil {
		return nil, err
	}

	if view == nil {
		return nil, nil
	}
	//	// TODO ： 校验当前用户是否有查看权限 缺失了，后续补充上

	roleRelView, _ := relRoleView.DaoGetter.GetByView(id)

	viewWithSourceBaseInfo := &ViewWithSourceBaseInfo{
		ViewWithSource: view,
	}
	if roleRelView != nil {
		viewWithSourceBaseInfo.Roles = roleRelView
	}

	return viewWithSourceBaseInfo, nil
}

func (s ServiceGetterImpl) GetViews(projectId int64, u *user.User) (*[]ViewBaseInfo, error) {
	projectDetail, _ := project.ServiceGetter.GetProjectDetail(projectId, u, false)

	fmt.Print(projectDetail)

	// todo 这里同样有判断project 权限问题，缺失的之后补上；

	views, err := DaoGetter.getViewBaseInfoByProject(projectId)

	if err != nil {
		return nil, err
	}

	return views, err
}

func NewServiceGetterImpl() *ServiceGetterImpl {
	return &ServiceGetterImpl{}
}

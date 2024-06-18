package project

import (
	"context"
	"davinci/module/relProjectAdmin"
	"davinci/module/relUserOrganization"
	"davinci/module/star"
	"davinci/module/user"
	"davinci/pkg/vars"
	"fmt"
)

var ServiceGetter Service

func init() {
	ServiceGetter = NewServiceGetterImpl()
}

type ServiceGetterImpl struct{}

func (this ServiceGetterImpl) GetProjectInfo(projectId int64, u *user.User) (ProjectInfo, error) {
	var projectInfo ProjectInfo
	detail, err := this.GetProjectDetail(projectId, u, false)

	fmt.Print("detail: ", detail)
	if err != nil {
		panic(err)
	}
	ConvertProjectDetailToProjectInfo(&projectInfo, detail)
	// 处理star

	hasStar, _ := star.DaoGetter.Select(u.Id, projectId, vars.STAR_TARGET_PROJECT)
	if hasStar != nil {
		projectInfo.ProjectWithCreateBy.IsStar = true
	}

	return projectInfo, nil
}

func (s ServiceGetterImpl) GetProjectDetail(projectId int64, u *user.User, modify bool) (*ProjectDetail, error) {
	projectDetail, err := DaoGetter.GetProjectDetail(projectId)

	fmt.Println("projectDetail: =======", projectDetail)

	if nil == projectDetail {
		message := fmt.Sprintf("Project %d not found", projectId)
		panic(message)
	}
	// 判断当前用户是否拥有查看当前项目详情的权限；

	rel, _ := relUserOrganization.DaoGetter.GetRel(u.Id, projectDetail.OrgId)
	fmt.Println(rel)

	relProAdmin, _ := relProjectAdmin.DaoGetter.GetByProjectAndUser(projectId, u.Id)

	isCreator := projectDetail.UserId == u.Id && !projectDetail.IsTransfer
	notOwner := !isCreator && relProAdmin == nil && (nil == rel || rel.Role != relUserOrganization.OwnerRole)

	if modify {
		//项目的创建人和当前项目对应组织的owner可以修改
		if notOwner {
			message := fmt.Sprintf("UnAuthorized, User %d does not have permission to modify project %d", u.Id, projectId)
			panic(message)
		}
	} else {
		//project所在org对普通成员project不可见
		if notOwner && projectDetail.Organization.MemberPermission < 1 && !projectDetail.Visibility {
			message := fmt.Sprintf("UnAuthorized, User %d have not permission to get project %d ", u.Id, projectId)
			panic(message)

		}
	}

	return projectDetail, err
}

func NewServiceGetterImpl() *ServiceGetterImpl {
	return &ServiceGetterImpl{}
}

func (s ServiceGetterImpl) GetFavoriteProjects(u *user.User) ([]ProjectWithCreateBy, error) {
	id := u.Id
	projects, err := DaoGetter.GetFavoriteProjects(id)
	println("projects: ", projects)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (s ServiceGetterImpl) GetProjects(ctx context.Context, u *user.User) ([]ProjectWithCreateBy, error) {
	id := u.Id
	projects, err := DaoGetter.GetProjectsByUser(id)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

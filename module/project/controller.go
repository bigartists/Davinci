package project

import (
	"davinci/module/role"
	"davinci/module/user"
	"davinci/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type ProjectController struct {
}

func NewProjectController() *ProjectController {
	return &ProjectController{}
}

func (this *ProjectController) GetProjects(c *gin.Context) {
	u := user.GetAuthUser(c)
	projects, err := ServiceGetter.GetProjects(c, u)

	if err != nil {
		c.JSON(500, utils.ResultWrapper(c)(nil, err.Error())(utils.Error))
		return
	}

	projectsDTO := ConvertProjectToDTO(projects)

	c.JSON(200, utils.ResultWrapper(c)(projectsDTO, "")(utils.OK))
}

func (this *ProjectController) GetFavoriteProjects(c *gin.Context) {
	u := user.GetAuthUser(c)
	projects, err := ServiceGetter.GetFavoriteProjects(u)

	if err != nil {
		c.JSON(500, utils.ResultWrapper(c)(nil, err.Error())(utils.Error))
		return
	}

	projectsDTO := ConvertProjectToDTO(projects)

	c.JSON(200, utils.ResultWrapper(c)(projectsDTO, "")(utils.OK))
}

func (this *ProjectController) GetRolesOfProject(c *gin.Context) {
	u := user.GetAuthUser(c)
	projectId, err := utils.GetInt64Param(c, "id")

	if projectId == 0 || err != nil {
		c.JSON(400, utils.ResultWrapper(c)(nil, "invalid project id")(utils.Error))
		return
	}

	roles, err := role.ServiceGetter.GetRolesByProjectId(projectId, u)
	if err != nil {
		c.JSON(500, utils.ResultWrapper(c)(nil, err.Error())(utils.Error))
		return
	}

	c.JSON(200, utils.ResultWrapper(c)(roles, "")(utils.OK))

}

func (this *ProjectController) GetProjectInfo(c *gin.Context) {
	id := cast.ToInt64(c.Param("id"))
	u := user.GetAuthUser(c)
	if id == 0 {
		c.JSON(400, utils.ResultWrapper(c)(nil, "invalid project id")(utils.Error))
		return
	}

	project, err := ServiceGetter.GetProjectInfo(id, u)
	if err != nil {
		c.JSON(500, utils.ResultWrapper(c)(nil, err.Error())(utils.Error))
		return
	}

	c.JSON(200, utils.ResultWrapper(c)(project, "")(utils.OK))
}

func (this *ProjectController) Build(r *gin.RouterGroup) {
	r.GET("/projects", this.GetProjects)
	r.GET("/projects/favorites", this.GetFavoriteProjects)
	r.GET("/projects/:id/roles", this.GetRolesOfProject)
	r.GET("/projects/:id", this.GetProjectInfo)
}

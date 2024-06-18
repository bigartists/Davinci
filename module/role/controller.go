package role

import (
	"github.com/gin-gonic/gin"
)

type RoleController struct {
}

func NewRoleController() *RoleController {
	return &RoleController{}
}

func (this *RoleController) GetRoles(c *gin.Context) {
	//u := middlewares.GetAuthUser(c)
	////projects, err := ServiceGetter.GetProjects(c, u)
	////roles, err := ServiceGetter.GetRolesByProjectId(c, u)
	//if err != nil {
	//	c.JSON(500, utils.ResultWrapper(c)(nil, err.Error())(utils.Error))
	//	return
	//}
	//
	//projectsDTO := ConvertProjectToDTO(projects)
	//
	//c.JSON(200, utils.ResultWrapper(c)(projectsDTO, "")(utils.OK))
}

func (this *RoleController) Build(r *gin.RouterGroup) {
	r.GET("/roles", this.GetRoles)
}

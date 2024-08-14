package organization

import (
	"davinci/module/user"
	"davinci/pkg/utils"
	"github.com/gin-gonic/gin"
)

type OrganizationController struct {
}

func NewOrganizationController() *OrganizationController {
	return &OrganizationController{}
}

func (this *OrganizationController) GetOrganizations(c *gin.Context) {
	u := user.GetAuthUser(c)
	organizations, err := ServiceGetter.GetOrganizations(c, u)

	if err != nil {
		c.JSON(500, utils.ResultWrapper(c)(nil, err.Error())(utils.Error))
		return
	}

	c.JSON(200, utils.ResultWrapper(c)(organizations, "")(utils.OK))

}

func (this *OrganizationController) GetOrgMembers(c *gin.Context) {
	orgId := c.Param("id")
	members, err := ServiceGetter.GetOrgMembers(orgId)
	if err != nil {
		c.JSON(500, utils.ResultWrapper(c)(nil, err.Error())(utils.Error))
		return
	}
	c.JSON(200, utils.ResultWrapper(c)(members, "")(utils.OK))
}

func (this *OrganizationController) Build(r *gin.RouterGroup) {
	r.GET("/organizations", this.GetOrganizations)
	r.GET("/organizations/:id/members", this.GetOrgMembers)
}

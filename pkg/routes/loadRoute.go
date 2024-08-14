package routes

import (
	"davinci/module/organization"
	"davinci/module/project"
	"davinci/module/qa"
	"davinci/module/source"
	"davinci/module/user"
	"davinci/module/view"
	"github.com/gin-gonic/gin"
)

func Build(r *gin.Engine) {
	group := r.Group("/api/v3") // *gin.RouterGroup
	user.NewUserController().Build(group)
	qa.NewQaController().Build(group)
	project.NewProjectController().Build(group)
	organization.NewOrganizationController().Build(group)
	source.NewSourceController().Build(group)
	view.NewViewController().Build(group)
}

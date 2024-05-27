package routes

import (
	"davinci/module/qa"
	"davinci/module/user"
	"github.com/gin-gonic/gin"
)

func Build(r *gin.Engine) {
	group := r.Group("/api/v1") // *gin.RouterGroup
	user.NewUserController().Build(group)
	qa.NewQaController().Build(group)
}

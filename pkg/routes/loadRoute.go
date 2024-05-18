package routes

import (
	"davinci/module/user"
	"github.com/gin-gonic/gin"
)

func Build(r *gin.Engine) {
	group := r.Group("/modi/v1") // *gin.RouterGroup
	user.NewAuthController().Build(group)
}

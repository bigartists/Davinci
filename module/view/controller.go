package view

import (
	"davinci/pkg/middlewares"
	"davinci/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type ViewController struct {
}

func NewViewController() *ViewController {
	return &ViewController{}
}

func (this *ViewController) GetViews(c *gin.Context) {
	u := middlewares.GetAuthUser(c)
	projectId := cast.ToInt64(c.DefaultQuery("projectId", "0"))

	fmt.Println("GetViews===", projectId)

	if projectId == 0 {
		c.JSON(400, utils.ResultWrapper(c)(nil, "projectId is required")(utils.Error))
		return
	}

	views, err := ServiceGetter.GetViews(projectId, u)

	if err != nil {
		c.JSON(500, utils.ResultWrapper(c)(nil, err.Error())(utils.Error))
		return
	}

	c.JSON(200, utils.ResultWrapper(c)(views, "")(utils.OK))
}

func (this *ViewController) GetViewDetail(c *gin.Context) {
	u := middlewares.GetAuthUser(c)
	id := cast.ToInt64(c.Param("id"))
	if id == 0 {
		c.JSON(400, utils.ResultWrapper(c)(nil, "viewid is required")(utils.Error))
		return
	}

	viewDetail, err := ServiceGetter.GetViewDetail(u, id)

	if err != nil {
		c.JSON(500, utils.ResultWrapper(c)(nil, err.Error())(utils.Error))
	}
	c.JSON(200, utils.ResultWrapper(c)(viewDetail, "")(utils.OK))
}

func (this *ViewController) Build(r *gin.RouterGroup) {
	r.GET("/views", this.GetViews)
	r.GET("/views/:id", this.GetViewDetail)
}

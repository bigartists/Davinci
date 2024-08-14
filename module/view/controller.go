package view

import (
	"davinci/module/user"
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
	u := user.GetAuthUser(c)
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
	u := user.GetAuthUser(c)
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

type QueryParams struct {
	Limit     int64         `form:"limit,default=500"`
	SourceId  int64         `form:"sourceId,default=0"`
	SQL       string        `form:"sql,default="`
	Variables []interface{} `form:"variables,default="`
}

func (this *ViewController) ExecuteSql(c *gin.Context) {
	u := user.GetAuthUser(c)

	var queryParams QueryParams
	fmt.Println("queryParams 66:", queryParams)
	if err := c.ShouldBind(&queryParams); err != nil {
		fmt.Println("Error binding params:", err)
		c.JSON(400, utils.ResultWrapper(c)(nil, err.Error())(utils.Error))
		return
	}
	fmt.Println("queryParams 72:", queryParams)

	rows, err := ServiceGetter.ExecuteSql(u, queryParams.SourceId, queryParams.SQL, queryParams.Variables, queryParams.Limit)
	if err != nil {
		c.JSON(500, utils.ResultWrapper(c)(nil, err.Error())(utils.Error))
		return
	}

	//ret := make(map[string]interface{}, 0)
	//ret["resultList"] = rows

	c.JSON(200, utils.ResultWrapper(c)(rows, "")(utils.OK))
}

func (this *ViewController) Build(r *gin.RouterGroup) {
	r.GET("/views", this.GetViews)
	r.GET("/views/:id", this.GetViewDetail)
	r.POST("/views/executesql", this.ExecuteSql)
}

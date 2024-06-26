package source

import (
	"davinci/pkg/middlewares"
	"davinci/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type SourceController struct {
}

func NewSourceController() *SourceController {
	return &SourceController{}
}

func (this *SourceController) GetSources(c *gin.Context) {
	u := middlewares.GetAuthUser(c)
	projectId := cast.ToInt64(c.DefaultQuery("projectId", "0"))

	sources, err := ServiceGetter.GetSources(projectId, u)

	if err != nil {
		c.JSON(500, utils.ResultWrapper(c)(nil, err.Error())(utils.Error))
		return
	}
	sourceDTO := ConvertSourcesToDTO(*sources)
	c.JSON(200, utils.ResultWrapper(c)(sourceDTO, "")(utils.OK))
}

func (this *SourceController) GetSourceDetail(c *gin.Context) {
	u := middlewares.GetAuthUser(c)
	id := cast.ToInt64(c.Param("id"))

	sourceDetail, err := ServiceGetter.GetSourceDetail(u, id)

	if err != nil {
		c.JSON(500, utils.ResultWrapper(c)(nil, err.Error())(utils.Error))
	}
	c.JSON(200, utils.ResultWrapper(c)(sourceDetail, "")(utils.OK))
}

func (this *SourceController) Build(r *gin.RouterGroup) {
	r.GET("/sources", this.GetSources)
	r.GET("/sources/:id", this.GetSourceDetail)
}

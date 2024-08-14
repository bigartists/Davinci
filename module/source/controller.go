package source

import (
	"davinci/module/user"
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
	u := user.GetAuthUser(c)
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
	u := user.GetAuthUser(c)
	id := cast.ToInt64(c.Param("id"))

	sourceDetail, err := ServiceGetter.GetSourceDetail(u, id)

	if err != nil {
		c.JSON(500, utils.ResultWrapper(c)(nil, err.Error())(utils.Error))
		return
	}
	c.JSON(200, utils.ResultWrapper(c)(sourceDetail, "")(utils.OK))
}

func (this *SourceController) GetSourceDatabases(c *gin.Context) {
	u := user.GetAuthUser(c)
	id := cast.ToInt64(c.Param("id"))

	databases, err := ServiceGetter.GetSourceDatabases(u, id)
	if err != nil {
		c.JSON(500, utils.ResultWrapper(c)(nil, err.Error())(utils.Error))
		return
	}

	res := map[string]interface{}{
		"sourceId":  id,
		"databases": databases,
	}

	c.JSON(200, utils.ResultWrapper(c)(res, "")(utils.OK))
}

func (this *SourceController) GetSourceTables(c *gin.Context) {
	u := user.GetAuthUser(c)
	id := cast.ToInt64(c.Param("id"))
	// http://localhost:5002/api/v3/sources/2/tables?dbName=analysis 获取dbName参数
	dbName := c.Query("dbName")
	tables, err := ServiceGetter.GetSourceTables(u, id, dbName)
	if err != nil {
		c.JSON(500, utils.ResultWrapper(c)(nil, err.Error())(utils.Error))
		return
	}
	// 遍历 tables  从 []string 变成  []map[string]interface{}{{"name": "table1", "type":"TABLE"}, {"name": "table2", "type":"TABLE"
	res := map[string]interface{}{
		"sourceId": id,
		"dbName":   dbName,
		"tables":   ConvertTablesToMap(tables),
	}
	c.JSON(200, utils.ResultWrapper(c)(res, "")(utils.OK))
}

func (this *SourceController) GetSourceTableColumns(c *gin.Context) {
	u := user.GetAuthUser(c)
	id := cast.ToInt64(c.Param("id"))
	dbName := c.Query("dbName")
	tableName := c.Query("tableName")

	columns, err := ServiceGetter.GetSourceTableColumns(u, id, dbName, tableName)
	if err != nil {
		c.JSON(500, utils.ResultWrapper(c)(nil, err.Error())(utils.Error))
	}
	res := map[string]interface{}{
		"sourceId":    id,
		"tableName":   tableName,
		"primaryKeys": []string{},
		"columns":     columns,
	}
	c.JSON(200, utils.ResultWrapper(c)(res, "")(utils.OK))
}

func (this *SourceController) Build(r *gin.RouterGroup) {
	r.GET("/sources", this.GetSources)
	r.GET("/sources/:id", this.GetSourceDetail)
	r.GET("/sources/:id/databases", this.GetSourceDatabases)
	r.GET("/sources/:id/tables", this.GetSourceTables)
	r.GET("/sources/:id/table/columns", this.GetSourceTableColumns)
}

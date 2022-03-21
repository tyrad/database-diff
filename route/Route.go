package route

import (
	"db-diff/route/api"
	"db-diff/route/site"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode("debug")
	// 网站配置
	site.ConfigWebSite(r)

	// 批量导入功能 文件上传
	r.POST("yxx/uploadFile", api.UploadFile)

	// 嵌套路由组  /shop/xx/oo
	dbApi := r.Group("/db")
	// 检测数据库可以连通
	dbApi.POST("/connectTest", api.DbConnectTest)
	dbApi.POST("/queryTable", api.DbQueryTable)
	dbApi.POST("/analyzeTables", api.AnalyzeTables)

	return r
}

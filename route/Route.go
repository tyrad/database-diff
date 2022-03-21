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

	// 关联查询字典值
	r.GET("/yxx/getHospitalCode", api.GetHospitalCode)

	// 加载区域编码
	r.GET("yxx/getAreaCode", api.GetAreaCode)

	// 批量导入功能 文件上传
	r.POST("yxx/uploadFile", api.UploadFile)

	// 嵌套路由组  /shop/xx/oo
	dbApi := r.Group("/db")
	dbApi.POST("/connectTest", api.DbConnectTest)

	return r
}

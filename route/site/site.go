package site

import (
	"db-diff/pkg/path"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ConfigWebSite(r *gin.Engine) {
	// 配置vue后台资源
	r.StaticFS("/manager", http.Dir(path.GenAppPath("runtime/public/dist")))
}

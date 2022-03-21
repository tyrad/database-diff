package api

import (
	"db-diff/model"
	"db-diff/pkg/app"
	"db-diff/pkg/db"
	"db-diff/pkg/e"
	"github.com/gin-gonic/gin"
)

func AnalyzeTables(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form model.DbAnalyzeTable
	)
	err := app.BindJSONAndValid(c, &form)
	if err != nil {
		appG.ResponseError(err)
		return
	}
	// 1. 检测是否可以
	enable, err := db.CheckConnectEnable(form.MainDb.ToDbConfig())
	if !enable {
		appG.ResponseMsgError(err, "主库连接错误")
		return
	}
	enable, err = db.CheckConnectEnable(form.CompareDb.ToDbConfig())
	if !enable {
		appG.ResponseMsgError(err, "对比库连接错误")
		return
	}

	// 2. 对表进行检测
	connect, err := db.Connect(form.MainDb.ToDbConfig())
	tables, err := db.QueryTables(connect, "", "", "")
	if err != nil {
		appG.ResponseError(err)
		return
	}
	appG.Response("成功", e.SUCCESS, tables)
}

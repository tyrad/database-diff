package api

import (
	"db-diff/model"
	"db-diff/pkg/app"
	"db-diff/pkg/db"
	"db-diff/pkg/e"
	"github.com/gin-gonic/gin"
)

func DbQueryTable(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form model.DbQueryTable
	)
	err := app.BindJSONAndValid(c, &form)
	if err != nil {
		appG.ResponseError(err)
		return
	}
	// 1. 检测是否可以
	enable, err := db.CheckConnectEnable(form.ToDbConfig())
	if !enable {
		appG.ResponseError(err)
		return
	}
	// 2. 查询表
	connect, err := db.Connect(form.ToDbConfig())
	tables, err := db.QueryTables(connect, form.FilterSchemaName, form.FilterTableOwner, form.FilterTableName)
	if err != nil {
		appG.ResponseError(err)
		return
	}
	appG.Response("成功", e.SUCCESS, tables)
}

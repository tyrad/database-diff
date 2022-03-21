package api

import (
	"db-diff/model"
	"db-diff/pkg/app"
	"db-diff/pkg/db"
	"db-diff/pkg/e"
	"fmt"
	"github.com/gin-gonic/gin"
)

func DbConnectTest(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form model.DbConfig
	)
	err := app.BindJSONAndValid(c, &form)
	if err != nil {
		appG.Response(fmt.Sprintf("%v", err), e.ERROR, nil)
		return
	}
	enable, err := db.CheckConnectEnable(form)
	fmt.Println("数据库连接是否可用:", enable, err)
	if enable {
		appG.Response("成功", e.SUCCESS, nil)
		return
	}
	appG.Response(fmt.Sprintf("%v", err), e.ERROR, nil)
}

func DbConnectTest2(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form model.DbConfig
	)
	err := app.BindAndValid(c, &form)
	if err != nil {
		appG.Response(fmt.Sprintf("%v", err), e.ERROR, nil)
		return
	}
	appG.Response(fmt.Sprintf("%v", form), e.SUCCESS, nil)
}

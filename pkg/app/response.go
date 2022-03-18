package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

/*
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}
*/

func (g *Gin) Response(msg string, errCode int, data interface{}) {
	// 方便前端抛出异常，这里统一状态码为 200 (业务正常返回数据)
	g.C.JSON(http.StatusOK, Response{
		Code: errCode,
		Msg:  msg,
		Data: data,
	})
	return
}

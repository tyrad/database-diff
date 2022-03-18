package app

import (
	"db-diff/pkg/e"
	"db-diff/pkg/logging"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusOK, e.INVALID_PARAMS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		logging.Error(err)
		return http.StatusInternalServerError, e.ERROR
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusOK, e.INVALID_PARAMS
	}

	return http.StatusOK, e.SUCCESS
}

func BindJSONAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.BindJSON(form)
	if err != nil {
		return http.StatusOK, e.INVALID_PARAMS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, e.ERROR
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusOK, e.INVALID_PARAMS
	}

	return http.StatusOK, e.SUCCESS
}

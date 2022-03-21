package app

import (
	"db-diff/pkg/e"
	"db-diff/pkg/logging"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) error {
	err := c.Bind(form)
	if err != nil {
		return err
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		logging.Error(err)
		return err
	}
	if !check {
		MarkErrors(valid.Errors)
		return e.MakeCommonErr("参数校验失败", nil, e.INVALID_PARAMS)
	}
	return nil
}

func BindJSONAndValid(c *gin.Context, form interface{}) error {
	err := c.BindJSON(form)
	if err != nil {
		return err
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return err
	}
	if !check {
		MarkErrors(valid.Errors)
		return e.MakeCommonErr("参数校验失败", nil, e.INVALID_PARAMS)
	}
	return nil
}

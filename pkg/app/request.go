package app

import (
	"db-diff/pkg/logging"
	"github.com/astaxie/beego/validation"
)

// MarkErrors 便于清空包中各处的log
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}
	return
}

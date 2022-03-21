package filter

import (
	"db-diff/model"
	"db-diff/pkg/db"
	"db-diff/pkg/e"
	"fmt"
)

type CheckConnectFilter struct {
	model model.DbAnalyzeTable
}

func (cf *CheckConnectFilter) Process(data Request) (Response, error) {
	mainDb, err := db.Connect(cf.model.MainDb.ToDbConfig())
	if err != nil {
		return nil, e.MakeCommonErr("主库连接失败", &err, e.ERROR)
	}
	compareDb, err := db.Connect(cf.model.CompareDb.ToDbConfig())
	if err != nil {
		return nil, e.MakeCommonErr("主库连接失败", &err, e.ERROR)
	}
	fmt.Println("验证数据库连接: 连接成功")
	return model.NewDbAnalyzeTableProcess(cf.model, mainDb, compareDb), nil
}

func NewCheckConnectFilter(analyzeTableModel model.DbAnalyzeTable) *CheckConnectFilter {
	return &CheckConnectFilter{model: analyzeTableModel}
}

package filter

import (
	"db-diff/model"
	"testing"
)

func TestChain_Process(t *testing.T) {

	// 1. 验证数据库连接
	connectFilter := NewCheckConnectFilter(model.DbAnalyzeTable{
		MainDb: model.DbAnalyzeTableItem{
			Host:     "10.67.78.133",
			User:     "phpg_nbphs_read",
			Password: "A4hPa8ECuUD7rSmFGNGFdmrU",
			DbName:   "chisapp",
			Port:     5432,
		},
		CompareDb: model.DbAnalyzeTableItem{
			Host:     "10.67.78.133",
			User:     "phpg_nbphs_read",
			Password: "A4hPa8ECuUD7rSmFGNGFdmrU",
			DbName:   "chisapp",
			Port:     5432,
		},
		TableNames: []string{"account"},
	})
	accountValid := NewChain("验证问题", connectFilter)
	_, err := accountValid.Process(0)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("验证成功")
}

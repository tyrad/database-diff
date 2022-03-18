package api

import (
	"db-diff/pkg/app"
	"db-diff/pkg/csv"
	"db-diff/pkg/e"
	"db-diff/pkg/path"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func GetAreaCode(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	code, err := getAllAreaCode()
	if err == nil {
		appG.Response("成功", e.SUCCESS, code)
		return
	}
	appG.Response(fmt.Sprintf("%v", err), e.ERROR, nil)
}

func getAllAreaCode() ([]map[string]string, error) {
	rows, err := csv.Read(path.GenAppPath(csv.AreaCodeCSV))
	if err == nil {
		log.Printf("数据读取成功 %v", rows)
	} else {
		return nil, e.MakeCommonErr(fmt.Sprintf("解析csv数据【%s】失败 %v", path.GenAppPath(csv.HospitalCsvCode), err), nil, e.ERROR)
	}
	var x []map[string]string
	for _, row := range rows {
		if len(row) < 3 {
			return nil, e.MakeCommonErr(fmt.Sprintf("数据缺失,至少有3列 【%v】", row), nil, e.ERROR)
		}
		x = append(x, map[string]string{
			"Name": row[0],
			"Code": row[1],
		})
	}
	return x, nil
}

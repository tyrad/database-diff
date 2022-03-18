package api

import (
	"db-diff/model"
	"db-diff/pkg/app"
	"db-diff/pkg/csv"
	"db-diff/pkg/e"
	"db-diff/pkg/path"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func GetHospitalCode(c *gin.Context) {
	appG := app.Gin{C: c}
	code := c.Query("name")
	right, err := readCsvPathRight(code)
	if err != nil {
		appG.Response(fmt.Sprintf("%v", err), e.ERROR, nil)
		return
	}
	appG.Response("成功", e.SUCCESS, right)
}

func readCsvPathRight(filterName string) (*[]model.HospitalCode, error) {
	res, err := csv.Read(path.GenAppPath(csv.HospitalCsvCode))
	if err == nil {
		log.Printf("数据读取成功 %v", res)
	} else {
		return nil, e.MakeCommonErr(fmt.Sprintf("解析csv数据【%s】失败 %v", path.GenAppPath(csv.HospitalCsvCode), err), nil, e.ERROR)
	}
	return toModels(res, filterName)
}

func toModels(rows [][]string, filterName string) (*[]model.HospitalCode, error) {
	var result []model.HospitalCode
	for _, item := range rows {
		t, err := toModel(item)
		if err != nil {
			return nil, e.MakeCommonErr(fmt.Sprintf("解析csv数据【%s】失败 %v", path.GenAppPath(csv.HospitalCsvCode), err), nil, e.ERROR)
		}
		if filterName != "" && strings.Contains(t.Name, filterName) {
			result = append(result, *t)
		}
		if filterName == "" {
			result = append(result, *t)
		}
	}
	return &result, nil
}

func toModel(row []string) (*model.HospitalCode, error) {
	if len(row) < 3 {
		return nil, e.MakeCommonErr(fmt.Sprintf("数据缺失,至少有3列 【%v】", row), nil, e.ERROR)
	}
	return &model.HospitalCode{
		Name: row[0],
		Code: row[2],
	}, nil
}

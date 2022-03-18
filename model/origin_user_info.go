package model

import (
	"fmt"
)

type OriginUserInfo struct {
	RegionName                  string
	UserName                    string
	Phone                       string
	Sex                         string
	GwAccount                   string
	RegionCode                  string
	InsertResult                string
	FixRegionAndGwAccountResult string
}

func MakeUserWithCsvRows(rows [][]string) (*[]OriginUserInfo, string) {
	var result []OriginUserInfo
	for _, item := range rows {
		t, err := MakeUserWithCsvRow(item)
		if err == "" {
			result = append(result, *t)
		}
		if err != "" {
			return nil, err
		}
	}
	return &result, ""
}

// MakeUserWithCsvRow 将数据转模型, 数据格式为: [ 大闫家卫生室,付永林,15662719300,男,zgzfyl ]
func MakeUserWithCsvRow(row []string) (*OriginUserInfo, string) {
	if len(row) < 5 {
		return nil, fmt.Sprintf("<span style='color:red'>数据异常,至少有5列 %v </span>", row)
	}
	// 所在卫生院名称,姓名,手机号,公卫账号
	// 简单验证有效性
	if row[0] == "" || row[1] == "" || row[2] == "" || row[3] == "" || row[4] == "" {
		return nil, fmt.Sprintf("<span style='color:red'>数据异常,至少有5列,每列数据均不能为空 %v </span>", row)
	}
	phone := row[2]
	if len(phone) != 11 {
		return nil, fmt.Sprintf("<span style='color:red'>数据异常,手机号需要为11位 %v </span>", row)
	}
	// 如果有第六列,则第六列为 RegionCode
	regionCode := ""
	insertResult := ""
	if len(row) >= 6 {
		regionCode = row[5]
	}
	if len(row) >= 7 {
		insertResult = row[6]
	}
	return &OriginUserInfo{
		RegionName:   row[0],
		UserName:     row[1],
		Phone:        row[2],
		Sex:          row[3],
		GwAccount:    row[4],
		RegionCode:   regionCode,
		InsertResult: insertResult,
	}, ""
}

package filter

import (
	"db-diff/model"
	"db-diff/pkg/e"
)

type CompareTable struct {
}

func (ct *CompareTable) Process(data Request) (Response, error) {
	md, ok := data.(*model.DbAnalyzeTableProcess)
	if !ok {
		return md, e.MakeCommonErr("数据强转失败 DbAnalyzeTableProcess", nil, e.ERROR)
	}
	//mainDb := md.MainDb
	//compareDb := md.CompareDB
	//tableNames := md.DbAnalyzeTable.TableNames

	// 根据名字获取table信息

	return nil, nil
}

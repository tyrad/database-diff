package filter

import (
	"db-diff/model"
	"db-diff/pkg/db"
	"db-diff/pkg/e"
	"fmt"
)

type CompareTable struct {
}

func NewCompareTable() *CompareTable {
	return &CompareTable{}
}

func (ct *CompareTable) Process(data Request) (Response, error) {
	md, ok := data.(*model.DbAnalyzeTableProcess)
	if !ok {
		return md, e.MakeCommonErr("数据强转失败 DbAnalyzeTableProcess", nil, e.ERROR)
	}
	mainDb := md.MainDb
	compareDb := md.CompareDB
	tableNames := md.DbAnalyzeTable.TableNames
	if len(tableNames) == 0 {
		return nil, e.MakeCommonErr("比较table数组为空", nil, e.ERROR)
	}
	// 根据名字获取table信息
	for _, tableName := range tableNames {
		fmt.Println("当前处理table", tableName)
		mainTableFields, mErr := db.DbAnalyzeTables(mainDb, tableName)
		compareTableFields, cErr := db.DbAnalyzeTables(compareDb, tableName)
		if cErr != nil {
			return nil, e.MakeCommonErr(fmt.Sprintf("查询主库【%s】表信息失败", tableName), nil, e.ERROR)
		}
		if mErr != nil {
			return nil, e.MakeCommonErr(fmt.Sprintf("查询对比库【%s】表信息失败", tableName), nil, e.ERROR)
		}
		fmt.Println("对比表字段信息:", tableName)
		// 1. 比较字段数量是否一致
		if len(mainTableFields) != len(compareTableFields) {
			fmt.Println(fmt.Sprintf("【%s】表字段不一致 -> TODO 点击查看详情", tableName))
			var mainTableFieldsNames []string
			var compareTableFieldsNames []string
			for _, ff := range mainTableFields {
				mainTableFieldsNames = append(mainTableFieldsNames, ff.ColumnName.String)
			}
			for _, ff := range compareTableFields {
				compareTableFieldsNames = append(compareTableFieldsNames, ff.ColumnName.String)
			}
			added, removed := Arrcmp(mainTableFieldsNames, compareTableFieldsNames)
			fmt.Println(fmt.Sprintf("概览【%s】表字段不一致。 对比主表新增字段%v  缺少字段%v", tableName, added, removed))
		}
		// 2. 对比字段是否一致
		for _, mFiled := range mainTableFields {
			columnName := mFiled.ColumnName
			// 查询出过滤信息
			var cFiled *model.DbTableInfo
			for _, ff := range compareTableFields {
				if ff.ColumnName == columnName {
					cFiled = ff
					break
				}
			}
			if cFiled == nil {
				fmt.Println(fmt.Sprintf("比较表缺少字段 %s", columnName.String))
				continue
			}
			// 比较具体信息
			isEqual, _, errDescription := mFiled.EqualAndExtractErrFields(cFiled)
			if isEqual {
				fmt.Println(fmt.Sprintf("【%s.%s】 检测通过", cFiled.TableName.String, cFiled.ColumnName.String))
			} else {
				fmt.Println(fmt.Sprintf("【%s】 检测失败，差异见 %v", cFiled.ColumnName.String, errDescription))
			}
		}
		// 3. 检测约束是否一致

	}
	return md, nil
}

func Arrcmp(src []string, dest []string) ([]string, []string) {
	msrc := make(map[string]byte) //按源数组建索引
	mall := make(map[string]byte) //源+目所有元素建索引
	var set []string              //交集
	//1.源数组建立map
	for _, v := range src {
		msrc[v] = 0
		mall[v] = 0
	}
	//2.目数组中，存不进去，即重复元素，所有存不进去的集合就是并集
	for _, v := range dest {
		l := len(mall)
		mall[v] = 1
		if l != len(mall) { //长度变化，即可以存
			l = len(mall)
		} else { //存不了，进并集
			set = append(set, v)
		}
	}
	//3.遍历交集，在并集中找，找到就从并集中删，删完后就是补集（即并-交=所有变化的元素）
	for _, v := range set {
		delete(mall, v)
	}
	//4.此时，mall是补集，所有元素去源中找，找到就是删除的，找不到的必定能在目数组中找到，即新加的
	var added, deleted []string
	for v, _ := range mall {
		_, exist := msrc[v]
		if exist {
			deleted = append(deleted, v)
		} else {
			added = append(added, v)
		}
	}
	return added, deleted
}

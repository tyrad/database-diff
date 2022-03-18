package csv

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

// 读取错误文件测试
func TestReadCsvPathError(t *testing.T) {
	res, err := Read("test2.csv")
	if err == nil {
		log.Printf("数据读取成功 %v", res)
	} else {
		log.Printf(err.Error())
	}
}

// 读取正确文件测试
func TestReadCsvPathRight(t *testing.T) {
	//res, err := Read("test.csv")
	//if err == nil {
	//	log.Printf("数据读取成功 %v", res)
	//} else {
	//	log.Printf(err.Error())
	//}
	//userList, err := model.MakeUserWithCsvRows(res)
	//if userList != nil {
	//	for _, item := range *userList {
	//		println(item.UserName)
	//	}
	//}
	//print(err)
}

func TestWrite(t *testing.T) {
	records := [][]string{
		{"first_name", "last_name", "occupation"},
		{"John", "Doe", "gardener"},
		{"Lucy", "Smith", "teacher"},
		{"Brian", "Bethamy", "programmer"},
	}

	err := Write(records, "output.csv")
	if err != nil {
		log.Println(err)
	}
}

func TestCC(t *testing.T) {
	text := "runtime/dataImport/import/2022-02-10-1644476188237/step0_原始数据.csv"
	index := strings.LastIndex(text, "/")
	if index >= 0 {
		content := text[0:index]
		fmt.Println(content)
	}
}

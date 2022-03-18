package csv

import (
	"db-diff/pkg/e"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

const HospitalCsvCode = "runtime/csv/nameToHospitalCode.csv"
const AreaCodeCSV = "runtime/csv/areaCodeAndUrl/areaCode.csv"

func Read(filePath string) ([][]string, error) {
	// 尝试打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return nil, e.MakeCommonErr(fmt.Sprintf("打开文件%s失败", filePath), &err, e.ERROR)
	}
	defer file.Close()

	// 读取csv文件
	r := csv.NewReader(file)
	r.FieldsPerRecord = -1

	// 遍历数据存到数组中
	var result [][]string
	for {
		// 逐行读取
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, record)
	}
	return result, nil
}

// Write 写文件 需要保证 filePath 存在
func Write(records [][]string, filePath string) error {

	f, err := os.Create(filePath)
	defer f.Close()

	if err != nil {
		fmt.Println(fmt.Sprintf("打开文件%s失败", filePath))
		return e.MakeCommonErr(fmt.Sprintf("打开文件%s失败", filePath), &err, e.ERROR)
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Println(fmt.Sprintf("- csv写入【%s】失败", filePath))
			return e.MakeCommonErr(fmt.Sprintf("csv写入【%s】失败", filePath), &err, e.ERROR)
		}
	}
	return nil
}

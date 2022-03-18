package gintemplate

import (
	"html/template"
	"time"
)

// Sub /*函数减法*/
func Sub(v1 int64, v2 int64) int64 {
	return v1 - v2
}

func Add(v1 int64, v2 int64) int64 {
	return v1 + v2
}

func AddInt(v1 int, v2 int) int {
	return v1 + v2
}

// IntRange /*遍历n-m*/
func IntRange(start, end int64) []int64 {
	n := end - start + 1
	result := make([]int64, n)
	var i int64
	for i = 0; i < n; i++ {
		result[i] = start + i
	}
	return result
}

func SafeHtml(html string) template.HTML {
	return template.HTML(html)
}

func DateFormat(sec int64, format string) string {
	return time.Unix(sec, 0).Format(format)
}

// str_time := time.Unix(1389058332, 0).Format("2006-01-02 15:04:05")

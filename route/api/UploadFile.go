package api

import (
	"db-diff/pkg/app"
	"db-diff/pkg/e"
	"db-diff/pkg/file"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"time"
)

const UploadFileFolder = "runtime/dataImport/import/"
const TYPE_ORG_FILE = "origin_data"
const TYPE_HOSPITAL_CODE_RELATION_FILE = "hospitalCodeRelation"

func UploadFile(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	f, header, err := c.Request.FormFile("file")
	unionFlag := c.Request.FormValue("unionFlag")
	uploadType := c.Request.FormValue("type")
	if f == nil || unionFlag == "" || err != nil {
		appG.Response("参数缺失 file 或  unionFlag 或 fileName", e.ERROR, nil)
		return
	}
	// 1. 文件存放到 runtime/dataImport/import/【日期 + unionFlag】  创建对应的目录
	timer := time.Now()
	curYear, curMonth, curDay := timer.Date()
	fmt.Printf("current date: %d-%02d-%02d\n", curYear, curMonth, curDay)

	folderName := fmt.Sprintf("%d-%02d-%02d-%s", curYear, curMonth, curDay, unionFlag)
	fmt.Printf(UploadFileFolder + folderName)
	mkdirFolder := UploadFileFolder + folderName

	err = file.IsNotExistMkDir(mkdirFolder)
	if err != nil {
		appG.Response("创建目录失败:"+mkdirFolder, e.ERROR, nil)
		return
	}

	// 2. 存放文件. 创建一个文件，文件名为filename，这里的返回值out也是一个File指针
	fileName := header.Filename
	if uploadType == TYPE_ORG_FILE {
		fileName = "step0_原始数据.csv"
	}
	if uploadType == TYPE_HOSPITAL_CODE_RELATION_FILE {
		fileName = "step0_卫生院编码对应关系.csv"
	}
	out, err := os.Create(mkdirFolder + "/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	// 将file的内容拷贝到out
	_, err = io.Copy(out, f)
	if err != nil {
		fmt.Sprintln(err)
		appG.Response(fmt.Sprintf("保存文件失败 %v", err), e.ERROR, nil)
		return
	}
	// 3. 成功后,返回文件路径
	appG.Response("成功", e.SUCCESS, mkdirFolder+"/"+fileName)
}

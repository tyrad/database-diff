package service

type DataImportStepOneCallback func(message string)

func DataImportStepOne(originFilePath string, hospitalCodeRelation string, callback DataImportStepOneCallback) {

}

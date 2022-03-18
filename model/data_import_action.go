package model

type DataImportActionModel struct {
	Step                 string `json:"step"`
	OriginFile           string `json:"originFile"`
	HospitalCodeRelation string `json:"hospitalCodeRelation"`
	Step1ResultCsv       string `json:"step1ResultCsv"`
	Step2ResultCsv       string `json:"Step2ResultCsv"`
}

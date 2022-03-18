package model

type DoctorData struct {
	DoctorInfo      *DoctorInfo
	DoctorLoginInfo *DoctorLoginModel
}

type DoctorInfo struct {
	ID                     int    `json:"id"`
	DoctorName             string `json:"doctorName"`
	HeadImageURL           string `json:"headImageUrl"`
	DoctorLoginID          int    `json:"doctorLoginId"`
	HospitalCode           string `json:"hospitalCode"`
	JobNumber              string `json:"jobNumber"`
	JobTitleCode           string `json:"jobTitleCode"`
	FirstClinicalSpecID    int    `json:"firstClinicalSpecId"`
	SecondClinicalSpecID   int    `json:"secondClinicalSpecId"`
	ExpertType             string `json:"expertType"`
	Intro                  string `json:"intro"`
	Expertise              string `json:"expertise"`
	Sex                    string `json:"sex"`
	FlagInvalid            int    `json:"flagInvalid"`
	CreateTime             string `json:"createTime"`
	UpdateTime             string `json:"updateTime"`
	AuthStatus             int    `json:"authStatus"`
	ExpertTypeName         string `json:"expertTypeName"`
	JobTitleName           string `json:"jobTitleName"`
	FirstDeptName          string `json:"firstDeptName"`
	SecondDeptName         string `json:"secondDeptName"`
	FirstClinicalSpecName  string `json:"firstClinicalSpecName"`
	SecondClinicalSpecName string `json:"secondClinicalSpecName"`
	IsHotSpec              int    `json:"isHotSpec"`
	HospitalName           string `json:"hospitalName"`
	HospitalLevel          string `json:"hospitalLevel"`
	PhoneNo                string `json:"phoneNo"`
	CurrentPage            int    `json:"currentPage"`
	PageSize               int    `json:"pageSize"`
	PageStart              int    `json:"pageStart"`
}

func NewDoctorInfo(name string, hospitalCode string, jobNumber string) *DoctorInfo {
	return &DoctorInfo{DoctorName: name, HospitalCode: hospitalCode, JobNumber: jobNumber}
}

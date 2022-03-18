package model

type DoctorLoginModel struct {
	ID              int         `json:"id"`
	DoctorName      string      `json:"doctorName"`
	PhoneNo         string      `json:"phoneNo"`
	LoginPassword   string      `json:"loginPassword"`
	MessagingID     interface{} `json:"messagingId"`
	HeadImageURL    string      `json:"headImageUrl"`
	FlagInvalid     int         `json:"flagInvalid"`
	CreateTime      string      `json:"createTime"`
	UpdateTime      string      `json:"updateTime"`
	CreateBy        string      `json:"createBy"`
	UpdateBy        string      `json:"updateBy"`
	RoleName        string      `json:"roleName"`
	RoleCode        string      `json:"roleCode"`
	StartTime       interface{} `json:"startTime"`
	EndTime         interface{} `json:"endTime"`
	DoctorID        int         `json:"doctorId"`
	MedicalWorkerID interface{} `json:"medicalWorkerId"`
}

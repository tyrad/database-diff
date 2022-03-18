package model

type DoctorCheckForm struct {
	Phone      string `form:"phone" valid:"Required;MaxSize(100)"`
	GwAccount  string `form:"gwAccount" valid:"MaxSize(255)"`
	GwPassword string `form:"gwPassword" valid:"MaxSize(255)"`
	AreaCode   string `form:"areaCode" valid:"MaxSize(255)"`
}

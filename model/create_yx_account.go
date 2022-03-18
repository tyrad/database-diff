package model

type CreateYxAccountForm struct {
	ID           string `form:"id" valid:"MaxSize(100)"`
	HospitalCode string `form:"hospitalCode" valid:"Required;MaxSize(100)"`
	GwAccount    string `form:"gwAccount" valid:"Required;MaxSize(255)"`
	Phone        string `form:"phone" valid:"Required;MaxSize(255)"`
	Sex          string `form:"sex" valid:"Required;MaxSize(255)"`
	UserName     string `form:"userName" valid:"Required;MaxSize(255)"`
}

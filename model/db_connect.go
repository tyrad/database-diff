package model

type DbConfig struct {
	Host     string `json:"host" form:"host" valid:"Required;MaxSize(100)"`
	User     string `json:"user" form:"user" valid:"Required;MaxSize(100)"`
	Password string `json:"password" form:"password" valid:"Required;MaxSize(100)"`
	DbName   string `json:"dbName" form:"dbName" valid:"Required;MaxSize(100)"`
	Port     int    `json:"port" form:"port" valid:"Required"`
}

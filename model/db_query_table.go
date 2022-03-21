package model

type DbQueryTable struct {
	Host             string `json:"host" valid:"Required;MaxSize(100)"`
	User             string `json:"user"  valid:"Required;MaxSize(100)"`
	Password         string `json:"password"  valid:"Required;MaxSize(100)"`
	DbName           string `json:"dbName" valid:"Required;MaxSize(100)"`
	Port             int    `json:"port" valid:"Required"`
	FilterTableName  string `json:"filterTableName"`
	FilterTableOwner string `json:"filterTableOwner"`
	FilterSchemaName string `json:"filterSchemaName"`
}

func (qt *DbQueryTable) ToDbConfig() *DbConfig {
	dbConfig := DbConfig{
		Host:     qt.Host,
		User:     qt.User,
		Password: qt.Password,
		DbName:   qt.DbName,
		Port:     qt.Port,
	}
	return &dbConfig
}

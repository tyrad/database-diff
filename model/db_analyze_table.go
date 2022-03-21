package model

type DbAnalyzeTable struct {
	MainDb     DbAnalyzeTableItem `json:"mainDb" valid:"Required;"`
	CompareDb  DbAnalyzeTableItem `json:"compareDb" valid:"Required;"`
	TableNames []string           `json:"tableNames" valid:"Required;MinSize(1)"`
}

type DbAnalyzeTableItem struct {
	Host     string `json:"host" valid:"Required;MaxSize(100)"`
	User     string `json:"user"  valid:"Required;MaxSize(100)"`
	Password string `json:"password"  valid:"Required;MaxSize(100)"`
	DbName   string `json:"dbName" valid:"Required;MaxSize(100)"`
	Port     int    `json:"port" valid:"Required"`
}

func (qt *DbAnalyzeTableItem) ToDbConfig() *DbConfig {
	dbConfig := DbConfig{
		Host:     qt.Host,
		User:     qt.User,
		Password: qt.Password,
		DbName:   qt.DbName,
		Port:     qt.Port,
	}
	return &dbConfig
}

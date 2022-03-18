package model

type DbConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"dbName"`
	Port     int    `json:"port"`
}

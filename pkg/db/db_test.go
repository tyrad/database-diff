package db

import (
	"db-diff/model"
	"fmt"
	_ "github.com/lib/pq"
	"testing"
)

var dbConfig = model.DbConfig{
	Host:     "10.67.78.133",
	User:     "phpg_nbphs_read",
	Password: "A4hPa8ECuUD7rSmFGNGFdmrU",
	DbName:   "chisapp",
	Port:     5432,
}

func TestConnect(t *testing.T) {
	enable, err := CheckConnectEnable(&dbConfig)
	fmt.Println("数据库连接是否可用:", enable, err)
	if err != nil {
		panic(err)
	}
}

func TestQueryOne(t *testing.T) {
	d2, err := Connect(&dbConfig)
	if err != nil {
		panic(err)
	}
	row := QueryRow(d2, "select * from phpg_nbphs.sys_config where config_name = $1", "enableMentalIllnessFaceReg")
	if row.Err() != nil {
		panic(row.Err())
	}
	fmt.Println(row)
}

type SysConfig struct {
	id          int64
	configName  string
	configValue string
}

func TestQueryOneMapToModel(t *testing.T) {
	d2, err := Connect(&dbConfig)
	if err != nil {
		panic(err)
	}
	row := QueryRow(d2, "select id, config_value, config_name from phpg_nbphs.sys_config where config_name = $1", "enableMentalIllnessFaceReg")
	if row.Err() != nil {
		panic(row.Err())
	}
	fmt.Println(row)
	var sysConfig SysConfig
	if err = row.Scan(
		&sysConfig.id, &sysConfig.configValue, &sysConfig.configName,
	); err != nil {
		panic(err)
	}
	fmt.Println(sysConfig)
}

func TestQueryMultiMapToModel(t *testing.T) {
	d2, err := Connect(&dbConfig)
	if err != nil {
		panic(err)
	}
	rows, err := Query(d2, "select id, config_value, config_name from phpg_nbphs.sys_config")
	if err != nil {
		panic(err)
	}
	var sysConfig SysConfig
	for rows.Next() {
		if err = rows.Scan(
			&sysConfig.id, &sysConfig.configValue, &sysConfig.configName); err != nil {
			panic(err)
		}
		fmt.Println(sysConfig)
	}
}

func TestQueryTables(t *testing.T) {
	d2, err := Connect(&dbConfig)
	if err != nil {
		panic(err)
	}

	tables, err := QueryTables(d2, "")
	if err != nil {
		panic(err)
	}
	for _, table := range tables {
		fmt.Println(table)
	}
}

func TestQueryTablesSchemaLimit(t *testing.T) {
	d2, err := Connect(&dbConfig)
	if err != nil {
		panic(err)
	}
	tables, err := QueryTables(d2, "phpg_nbphs")
	if err != nil {
		panic(err)
	}
	for _, table := range tables {
		fmt.Println(table)
	}
}

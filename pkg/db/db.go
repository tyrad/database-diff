package db

import (
	"database/sql"
	"db-diff/model"
	"fmt"
	_ "github.com/lib/pq"
)

// Connect 连接数据库
func Connect(config *model.DbConfig) (db *sql.DB, err error) {
	connect := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.User, config.Password, config.Host, config.Port, config.DbName)
	// Open只是验证参数是否合法，而不会创建和数据库的连接
	return sql.Open("postgres", connect)
}

// CheckConnectEnable 检查是否连通
func CheckConnectEnable(config *model.DbConfig) (bool, error) {
	connect, err := Connect(config)
	if err != nil {
		return false, err
	}
	err = connect.Ping()
	if err != nil {
		return false, err
	}
	return true, nil
}

// QueryRow 查询一行数据
func QueryRow(db *sql.DB, sql string, args ...interface{}) *sql.Row {
	row := db.QueryRow(sql, args...)
	return row
}

// Query 支持多条
func Query(db *sql.DB, sql string, args ...interface{}) (*sql.Rows, error) {
	return db.Query(sql, args...)
}

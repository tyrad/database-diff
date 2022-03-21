package db

import (
	"database/sql"
	"db-diff/model"
)

func DbAnalyzeTables(db *sql.DB, schema string, owner string, tableName string) ([]*model.DbTable, error) {

	return nil, nil
}

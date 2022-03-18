package db

import (
	"database/sql"
	"db-diff/model"
)

func QueryTables(db *sql.DB, schema string) ([]*model.DbTable, error) {
	var rows *sql.Rows
	var err error
	baseSql := "select schemaname, tablename, tableowner, hasindexes from pg_tables"
	if schema != "" {
		baseSql += " where schemaname = $1"
		rows, err = Query(db, baseSql, schema)
	} else {
		rows, err = Query(db, baseSql)
	}
	if err != nil {
		panic(err)
	}

	var tables []*model.DbTable
	for rows.Next() {
		var table model.DbTable
		err = rows.Scan(
			&table.SchemaName, &table.TableName, &table.TableOwner, &table.HasIndexes)
		if err != nil {
			return nil, err
		}
		tables = append(tables, &table)
	}
	return tables, nil
}

package db

import (
	"database/sql"
	"db-diff/model"
	"db-diff/pkg/logging"
)

func QueryTables(db *sql.DB, schema string, owner string, tableName string) ([]*model.DbTable, error) {
	var rows *sql.Rows
	var err error
	baseSql := `
	select schemaname, t.tablename, array_to_string(array_agg(tableowner), ',')
	from (
			 select schemaname, tablename, tableowner
			 from pg_tables
			 where schemaname like '%' || $1 || '%' and tablename like '%' || $2 || '%' and tableowner like '%' || $3 || '%'
		 ) t 
	group by tablename, schemaname`
	logging.Info(baseSql)
	rows, err = Query(db, baseSql, schema, tableName, owner)
	if err != nil {
		return nil, err
	}

	var tables []*model.DbTable
	for rows.Next() {
		var table model.DbTable
		err = rows.Scan(
			&table.SchemaName, &table.TableName, &table.TableOwner)
		if err != nil {
			return nil, err
		}
		tables = append(tables, &table)
	}
	return tables, nil
}

package db

import (
	"database/sql"
	"db-diff/model"
)

func DbAnalyzeConstraint(db *sql.DB, tableName string) ([]*model.DbTableConstraint, error) {
	rows, err := Query(db, `
		SELECT rel.relname                   table_name,
			   nsp.nspname                   schema_name,
			   con.contype                   con_type,
			   conname                       con_name,
			   pg_get_constraintdef(con.oid) constrait_description
		FROM pg_catalog.pg_constraint con
				 INNER JOIN pg_catalog.pg_class rel ON rel.oid = con.conrelid
				 INNER JOIN pg_catalog.pg_namespace nsp ON nsp.oid = connamespace
		WHERE rel.relname = $1
`, tableName)
	if err != nil {
		return nil, err
	}
	var tables []*model.DbTableConstraint
	for rows.Next() {
		var tableInfo model.DbTableConstraint
		if err = rows.Scan(
			&tableInfo.TableName,
			&tableInfo.SchemaName,
			&tableInfo.ConType,
			&tableInfo.ConName,
			&tableInfo.ConstraintDescription,
		); err != nil {
			return nil, err
		}
		tables = append(tables, &tableInfo)
	}
	return tables, nil
}

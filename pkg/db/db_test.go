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

	tables, err := QueryTables(d2, "", "", "")
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
	tables, err := QueryTables(d2, "phpg_nbphs", "", "")
	if err != nil {
		panic(err)
	}
	for _, table := range tables {
		fmt.Println(table)
	}
}

func TestGetTablesInfo(*testing.T) {
	d2, err := Connect(&dbConfig)
	if err != nil {
		panic(err)
	}
	tables, err := DbAnalyzeTables(d2, "sys_config")
	if err != nil {
		panic(err)
	}
	fmt.Println(tables)
}

func TestGetTablesInfo2(*testing.T) {
	d2, err := Connect(&dbConfig)
	if err != nil {
		panic(err)
	}
	tables, err := DbAnalyzeConstraint(d2, "sys_config")
	if err != nil {
		panic(err)
	}
	fmt.Println(tables)
}

func TestGetTableInfo(*testing.T) {
	d2, err := Connect(&dbConfig)
	if err != nil {
		panic(err)
	}
	rows, err := Query(d2, `
SELECT table_catalog /* 当前数据库名称 */,
       table_schema /* schema名称*/,
       table_name /* 表名 */,
       column_name /* 列表(字段名 )*/,
       ordinal_position /* 字段排序，从1开始 */,
       column_default /*字段默认值的表达式，例如自增序列: nextval(personal_id_seq'::regclass)  */,
       is_nullable /*描述是否可以为空。设置字段为空并不是设置字段不可为空的唯一方法 */,
       data_type /*数据类型   timestamp分为有时区、无时区。 TIMESTAMP/TIMESTAMPTZ   timestamp without time zone */,
       character_maximum_length /*字符串最大长度*/,
       character_octet_length /*标示字符串类型(character or bit string)，数据最大字节长度。非字符串(character or bit string)类型为null。和character_maximum_length的长度、服务器编码有关*/,
       numeric_precision /* numeric 类型，有效数字的数量 ,例如smallint=16 bigint=64 */,
       numeric_scale, /*小数有效数字*/
       datetime_precision /* 描述date、timestamp、interval type的精度。即秒值后面的小数位数。   date 0  timestamp=6     */,
       interval_type,
       interval_precision,
       udt_name, /*字段类型 名称*/
       (
           SELECT pg_catalog.col_description(c.oid, columns.ordinal_position::int)
           FROM pg_catalog.pg_class c
           WHERE c.oid = (SELECT ('"' || columns.table_name || '"')::regclass::oid)
             AND c.relname = columns.table_name
       ) AS column_comment
FROM information_schema.columns
WHERE table_name = 'hd_personal_base_others'
order by ordinal_position
`)
	if err != nil {
		panic(err)
	}
	var tables []*model.DbTableInfo
	for rows.Next() {
		var tableInfo model.DbTableInfo
		if err = rows.Scan(
			&tableInfo.TableCatalog,
			&tableInfo.TableSchema,
			&tableInfo.TableName,
			&tableInfo.ColumnName,
			&tableInfo.OrdinalPosition,
			&tableInfo.ColumnDefault,
			&tableInfo.IsNullable,
			&tableInfo.DataType,
			&tableInfo.CharacterMaximumLength,
			&tableInfo.CharacterOctetLength,
			&tableInfo.NumericPrecision,
			&tableInfo.NumericScale,
			&tableInfo.DatetimePrecision,
			&tableInfo.IntervalType,
			&tableInfo.IntervalPrecision,
			&tableInfo.UdtName,
			&tableInfo.ColumnComment,
		); err != nil {
			panic(err)
		}
		fmt.Println(tableInfo)
		_ = append(tables, &tableInfo)
	}
	fmt.Println(tables)
}

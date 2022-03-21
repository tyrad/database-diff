package model

import "database/sql"

type DbTableInfo struct {
	TableCatalog           sql.NullString `json:"tableCatalog"`           //  当前数据库名称
	TableSchema            sql.NullString `json:"tableSchema"`            // schema名称
	TableName              sql.NullString `json:"tableName"`              // 表名
	ColumnName             sql.NullString `json:"columnName"`             // 列表(字段名 )
	OrdinalPosition        sql.NullInt16  `json:"ordinalPosition"`        // 字段排序，从1开始
	ColumnDefault          sql.NullString `json:"columnDefault"`          // 字段默认值的表达式，例如自增序列: nextval(personal_id_seq'::regclass)
	IsNullable             sql.NullString `json:"isNullable"`             // 描述是否可以为空。设置字段为空并不是设置字段不可为空的唯一方法
	DataType               sql.NullString `json:"dataType"`               // 数据类型   timestamp分为有时区、无时区。 TIMESTAMP/TIMESTAMPTZ   timestamp without time zone
	CharacterMaximumLength sql.NullInt64  `json:"characterMaximumLength"` // 字符串最大长度
	CharacterOctetLength   sql.NullInt64  `json:"characterOctetLength"`   // 标示字符串类型(character or bit string)，数据最大字节长度。非字符串(character or bit string)类型为null。和character_maximum_length的长度、服务器编码有关
	NumericPrecision       sql.NullString `json:"numericPrecision"`       // numeric 类型，有效数字的数量  例如smallint=16 bigint=64
	NumericScale           sql.NullString `json:"numericScale"`           // 小数有效数字
	DatetimePrecision      sql.NullInt64  `json:"datetimePrecision"`      //  描述date、timestamp、interval type的精度。即秒值后面的小数位数。   date 0  timestamp=6
	IntervalType           sql.NullString `json:"intervalType"`
	IntervalPrecision      sql.NullString `json:"intervalPrecision"`
	UdtName                sql.NullString `json:"udtName"`       // 字段类型 名称*
	ColumnComment          sql.NullString `json:"columnComment"` // 字段备注
}

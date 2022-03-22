package service

import (
	"database/sql"
	"db-diff/model"
	"fmt"
	"testing"
)

func TestReflect(*testing.T) {

	typeOfHero := model.DbTableInfo{
		TableCatalog: sql.NullString{
			String: "11",
			Valid:  true,
		},
		TableSchema:            sql.NullString{},
		TableName:              sql.NullString{},
		ColumnName:             sql.NullString{},
		OrdinalPosition:        sql.NullInt16{},
		ColumnDefault:          sql.NullString{},
		IsNullable:             sql.NullString{},
		DataType:               sql.NullString{},
		CharacterMaximumLength: sql.NullInt64{},
		CharacterOctetLength:   sql.NullInt64{},
		NumericPrecision:       sql.NullString{},
		NumericScale:           sql.NullString{},
		DatetimePrecision:      sql.NullInt64{},
		IntervalType:           sql.NullString{},
		IntervalPrecision:      sql.NullString{},
		UdtName:                sql.NullString{},
		ColumnComment:          sql.NullString{},
	}

	typeOfHero2 := model.DbTableInfo{
		TableCatalog: sql.NullString{
			String: "11",
			Valid:  true,
		},
		TableSchema: sql.NullString{},
		TableName: sql.NullString{
			String: "gogs",
			Valid:  true,
		},
		ColumnName:      sql.NullString{},
		OrdinalPosition: sql.NullInt16{},
		ColumnDefault: sql.NullString{
			String: "gogs",
			Valid:  true,
		},
		IsNullable:             sql.NullString{},
		DataType:               sql.NullString{},
		CharacterMaximumLength: sql.NullInt64{},
		CharacterOctetLength:   sql.NullInt64{},
		NumericPrecision:       sql.NullString{},
		NumericScale:           sql.NullString{},
		DatetimePrecision:      sql.NullInt64{},
		IntervalType:           sql.NullString{},
		IntervalPrecision:      sql.NullString{},
		UdtName:                sql.NullString{},
		ColumnComment:          sql.NullString{},
	}

	equal := typeOfHero.Equal(&typeOfHero2)
	fmt.Println(equal)

	isEqu, fields, description := typeOfHero.EqualAndExtractErrFields(&typeOfHero2)
	fmt.Println(isEqu, fields, description)
}

package model

import (
	"database/sql"
	"fmt"
	"reflect"
)

type DbTableConstraint struct {
	TableName             sql.NullString `json:"tableName"`             // 表名
	SchemaName            sql.NullString `json:"schemaName"`            // schema名
	ConType               sql.NullString `json:"conType"`               // eg: p u ..
	ConName               sql.NullString `json:"conName"`               // sign_pay_record_pk
	ConstraintDescription sql.NullString `json:"constraintDescription"` // PRIMARY KEY (id)
}

// Equal 比较是否相同
func (f *DbTableConstraint) Equal(other *DbTableConstraint) bool {
	val := reflect.ValueOf(f).Elem()
	otherFields := reflect.Indirect(reflect.ValueOf(other))
	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		value := val.Field(i)
		otherValue := otherFields.FieldByName(typeField.Name)
		if value.Interface() != otherValue.Interface() {
			return false
		}
	}
	return true
}

// EqualAndExtractErrFields 比较字段是否相同，并提取不同的字段
func (f *DbTableConstraint) EqualAndExtractErrFields(other *DbTableConstraint) (bool, []string, []string) {
	val := reflect.ValueOf(f).Elem()
	otherFields := reflect.Indirect(reflect.ValueOf(other))
	var errFields []string
	var descriptionField []string
	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		value := val.Field(i)
		otherValue := otherFields.FieldByName(typeField.Name)
		if value.Interface() != otherValue.Interface() {
			errFields = append(errFields, typeField.Name)
			s := FieldComment[typeField.Name]
			if s == "" {
				s = typeField.Name
			}
			descriptionField = append(descriptionField, fmt.Sprintf("%s【%s】", s, typeField.Name))
		}
	}
	if len(errFields) > 0 {
		return false, errFields, descriptionField
	}
	return true, errFields, descriptionField
}

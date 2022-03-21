package model

type DbTable struct {
	SchemaName string `json:"schemaName"`
	TableName  string `json:"tableName"`
	TableOwner string `json:"tableOwner"`
}

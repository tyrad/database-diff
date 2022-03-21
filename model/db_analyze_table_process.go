package model

import "database/sql"

type DbAnalyzeTableProcess struct {
	DbAnalyzeTable *DbAnalyzeTable
	MainDb         *sql.DB
	CompareDB      *sql.DB
}

func NewDbAnalyzeTableProcess(dbAnalyzeTable DbAnalyzeTable, mainDb *sql.DB, compareDB *sql.DB) *DbAnalyzeTableProcess {
	return &DbAnalyzeTableProcess{DbAnalyzeTable: &dbAnalyzeTable, MainDb: mainDb, CompareDB: compareDB}
}

package database

import "database/sql"

func OpenDB() (*sql.DB, error) {
	return sql.Open("mysql", "root:jackY^45@tcp(localhost:3306)/latihan_db?parseTime=true")
}

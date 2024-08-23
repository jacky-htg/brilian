package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func openDB() (*sql.DB, error) {
	return sql.Open("mysql", "root:1234@tcp(localhost:3306)/latihan_db?parseTime=true")
}

type Instructor struct {
	Id   int
	Name string
	Loc  string
}

func findInstructor(db *sql.DB) error {
	var instructor Instructor
	const q = `SELECT id, name,loc FROM instructors WHERE id=1`
	err := db.QueryRow(q).Scan(&instructor.Id, &instructor.Name, &instructor.Loc)
	if err != nil {
		return err
	}
	fmt.Println(instructor)
	return nil
}

func main() {
	db, err := openDB()
	if err != nil {
		log.Fatalf("error: connecting to db: %s", err)
	}
	defer db.Close()

	err = findInstructor(db)
	if err != nil {
		println(err.Error())
	}
}

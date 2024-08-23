package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User struct
type Instructor struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:100"`
	// Email string `gorm:"uniqueIndex;size:100"`
}

func main() {
	// Data Source Name (DSN)
	dsn := "root:1234@tcp(127.0.0.1:3306)/latihan_db?charset=utf8mb4&parseTime=True&loc=Local"

	// Koneksi ke database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	// Migrasi schema
	db.AutoMigrate(&Instructor{})

	// Create
	//user := Instructor{Name: "John Doe"} //, Email: "john@example.com"}
	//db.Create(&user)

	// Read
	var readUser Instructor
	db.First(&readUser, 6) // Cari user dengan ID 1
	fmt.Println("Read User:", readUser)

	// Update
	// db.Model(&readUser).Update("Email", "john.doe@example.com")

	// Delete
	//db.Delete(&readUser)
}

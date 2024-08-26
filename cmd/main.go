package main

import (
	"log"

	"github.com/GuiaBolso/darwin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jacky-htg/brilian/pkg/database"
)

var (
	migrations = []darwin.Migration{
		{
			Version:     1,
			Description: "Creating table users",
			Script: `CREATE TABLE users (
						id INT NOT NULL auto_increment, 
						name 		VARCHAR(100) NOT NULL,
						password CHAR(36) NOT NULL,
						email VARCHAR(255) NOT NULL UNIQUE,
						created_at timestamp NOT NULL default NOW(),
						created_by int NOT NULL,
						updated_at timestamp NOT NULL default NOW(),
						updated_by int NOT NULL,  
						PRIMARY KEY (id)
					 ) ENGINE=InnoDB CHARACTER SET=utf8;`,
		},
	}
)

func main() {
	db, err := database.OpenDB()

	if err != nil {
		log.Fatal(err)
	}

	driver := darwin.NewGenericDriver(db, darwin.MySQLDialect{})

	d := darwin.New(driver, migrations, nil)
	err = d.Migrate()

	if err != nil {
		log.Println(err)
	}
}

package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jacky-htg/brilian/pkg/database"
	"github.com/jacky-htg/brilian/route"
)

func main() {

	logErr := log.New(os.Stdout, "ERROR : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	//info := log.New(os.Stdout, "INFO : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	db, err := database.OpenDB()
	if err != nil {
		log.Fatalf("error: connecting to db: %s", err)
	}
	defer db.Close()

	logErr.Fatal(http.ListenAndServe(":9000", route.InitRoute(db, logErr)))
}

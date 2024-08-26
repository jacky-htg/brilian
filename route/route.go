package route

import (
	"database/sql"
	"log"

	"github.com/jacky-htg/brilian/handlers"
	"github.com/julienschmidt/httprouter"
)

func InitRoute(db *sql.DB, log *log.Logger) *httprouter.Router {
	router := httprouter.New()

	user := handlers.UserHandler{Db: db, Log: log}

	router.GET("/users/:id", user.Get)
	router.POST("/users", user.Create)
	router.PUT("/users/:id", user.Update)
	router.DELETE("/users/:id", user.Delete)
	return router

}

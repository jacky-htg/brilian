package route

import (
	"database/sql"
	"log"

	"github.com/jacky-htg/brilian/handlers"
	"github.com/jacky-htg/brilian/middleware"
	"github.com/julienschmidt/httprouter"
)

func InitRoute(db *sql.DB, log *log.Logger) *httprouter.Router {
	router := httprouter.New()

	user := handlers.UserHandler{Db: db, Log: log}

	mid := middleware.Middleware{DB: db, Log: log}
	middlewares := []middleware.MiddHandler{
		mid.AuthorizationMiddleware, // Apply the token check middleware
	}

	router.GET("/users", user.List)
	router.GET("/users/:id", mid.InitMiddleware(middlewares, user.Get, "/users/:id"))
	router.POST("/users", mid.InitMiddleware(middlewares, user.Create, "/users"))
	router.PUT("/users/:id", user.Update)
	router.DELETE("/users/:id", user.Delete)
	router.POST("/login", user.Login)
	return router

}

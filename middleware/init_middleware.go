package middleware

import (
	"database/sql"
	"log"

	"github.com/julienschmidt/httprouter"
)

type Middleware struct {
	DB  *sql.DB
	Log *log.Logger
}

type MiddHandler func(httprouter.Handle, string) httprouter.Handle

func (u *Middleware) InitMiddleware(mw []MiddHandler, handler httprouter.Handle, path string) httprouter.Handle {
	for i := len(mw) - 1; i >= 0; i-- {
		h := mw[i]
		if h != nil {
			handler = h(handler, path)
		}
	}

	return handler
}

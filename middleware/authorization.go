package middleware

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (u *Middleware) AuthorizationMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id, _ := strconv.Atoi(ps.ByName("id"))
		println(r.Method, r.URL.Path)
		if !isAuthorization(id) {
			http.Error(w, "Invalid user", http.StatusUnauthorized)
			return
		}

		// Proceed to the next handler if the token is valid
		next(w, r, ps)
	}
}

func isAuthorization(id int) bool {
	// Implement your token validation logic here (e.g., check against a database or a secret
	return id == 1
}

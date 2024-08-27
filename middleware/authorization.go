package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/jacky-htg/brilian/models"
	"github.com/jacky-htg/brilian/pkg/token"
	"github.com/jacky-htg/brilian/repositories"
	"github.com/julienschmidt/httprouter"
)

func (u *Middleware) AuthorizationMiddleware(next httprouter.Handle, path string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Invalid Authorization header", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if !u.IsAuthorization(r.Context(), token) {
			http.Error(w, "Invalid user", http.StatusUnauthorized)
			return
		}

		// Proceed to the next handler if the token is valid
		next(w, r, ps)
	}
}

func (u *Middleware) IsAuthorization(ctx context.Context, mytoken string) bool {
	isValid, email := token.ValidateToken(mytoken)
	if !isValid {
		return false
	}

	userRepo := repositories.UserRepository{Db: u.DB, Log: u.Log, UserEntity: models.User{Email: email}}
	err := userRepo.GetByEmail(ctx)
	return err == nil
}

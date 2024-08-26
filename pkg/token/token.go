package token

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// MyCustomClaims struct
type MyCustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

var mySigningKey = []byte(os.Getenv("TOKEN_SALT"))

// ValidateToken for check token validation
func ValidateToken(myToken string) (bool, string) {
	token, err := jwt.ParseWithClaims(myToken, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySigningKey), nil
	})

	if err != nil {
		return false, ""
	}

	claims := token.Claims.(*MyCustomClaims)
	return token.Valid, claims.Email
}

// ClaimToken function
func ClaimToken(email string) (string, error) {
	claims := MyCustomClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with our secret
	return token.SignedString(mySigningKey)
}

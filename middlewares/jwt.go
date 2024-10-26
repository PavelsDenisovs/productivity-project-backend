package middlewares

import (
	"net/http"
	"strings"
  "os"
	"context"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

    if err != nil {
      http.Error(w, "Invalid token", http.StatusUnauthorized)
      return
    }

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := uint(claims["userID"].(float64))
			ctx := context.WithValue(r.Context(), "userID", userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
	})
}
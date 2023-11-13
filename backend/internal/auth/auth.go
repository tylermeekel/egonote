package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string
	jwt.RegisteredClaims
}

func SignJWT(username string, expTime time.Time) (string, error) {
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func parseJWT(tokenString string) (string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.Username, nil
	} else {
		return "", errors.New("invalid token")
	}
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt_token")
		if err != nil{
			ctx := context.WithValue(r.Context(), "username", "")
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		jwt := cookie.Value
		username, err := parseJWT(jwt)
		if err != nil{
			ctx := context.WithValue(r.Context(), "username", "")
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		} else{
			ctx := context.WithValue(r.Context(), "username", username)
			fmt.Println(username)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
	})
}

func getUsernameFromContext(r *http.Request) string {
	username := r.Context().Value("username").(string)
	return username
}

package auth

import (
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	ID int
	jwt.RegisteredClaims
}

func SignJWT(id int, expTime time.Time) (string, error) {
	claims := &Claims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func parseJWT(tokenString string) (int, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.ID, nil
	} else {
		return 0, errors.New("invalid token")
	}
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt_token")
		if err != nil {
			ctx := context.WithValue(r.Context(), "userID", -1)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		jwt := cookie.Value
		userID, err := parseJWT(jwt)
		if err != nil {
			ctx := context.WithValue(r.Context(), "userID", -1)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		} else {
			ctx := context.WithValue(r.Context(), "userID", userID)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
	})
}

func GetUserIDFromContext(r *http.Request) int {
	userID := r.Context().Value("userID").(int)
	return userID
}

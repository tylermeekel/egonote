package routes

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/tylermeekel/egonote/internal/auth"
	"github.com/tylermeekel/egonote/internal/data"
)

// InitServerMux creates and returns a *chi.Mux
// with all of the currently implemented routes
func InitServerMux(db data.Database) *chi.Mux {
	m := chi.NewMux()

	m.Use(auth.AuthMiddleware)
	m.Use(cors.Handler(
		cors.Options{
			AllowedOrigins: []string{"https://*", "http://*"},
			// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300,
		},
	))

	noteRouter := NoteRouter{DB: db}
	m.Mount("/notes", noteRouter.Routes())

	userRouter := UserRouter{DB: db}
	m.Mount("/users", userRouter.Routes())

	return m
}

func GetIDParam(r *http.Request) int {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return -1
	}
	return id
}

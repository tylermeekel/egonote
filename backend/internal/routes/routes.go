package routes

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/tylermeekel/egonote/internal/data"
)

type Router interface{
	Routes() *chi.Mux
}

// InitServerMux creates and returns a *chi.Mux
// with all of the currently implemented routes
func InitServerMux(db data.Database) *chi.Mux {
	m := chi.NewMux()

	noteRouter := NoteRouter{DB: db}
	m.Mount("/notes", noteRouter.Routes())

	return m
}

func GetIDParam(r *http.Request) int{
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil{
		return -1
	}
	return id
}

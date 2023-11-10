package routes

import (
	"github.com/go-chi/chi/v5"
)

// InitServerMux creates and returns a *chi.Mux
// with all of the currently implemented routes
func InitServerMux() *chi.Mux {
	m := chi.NewMux()

	m.Mount("/notes", noteRouter())

	return m
}

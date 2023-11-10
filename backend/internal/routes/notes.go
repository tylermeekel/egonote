package routes

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/tylermeekel/egonote/internal/data"
	"github.com/tylermeekel/egonote/internal/utils"
)

type NoteRouter struct {
	DB data.Database
}

// noteRouter creates and returns a router for
// accessing the notes API
func (n *NoteRouter) Routes() *chi.Mux {
	r := chi.NewMux()

	r.Get("/{id}", n.getNote)

	return r
}

// getNote gets a note and returns it as JSON
func (n *NoteRouter) getNote(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Write([]byte("Error"))
	}

	note := n.DB.GetNote(id)
	w.Header().Set("Content-Type", "application/json")
	utils.WriteJSON(w, note)
}

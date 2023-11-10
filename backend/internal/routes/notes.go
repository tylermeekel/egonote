package routes

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/tylermeekel/egonote/internal/types"
	"github.com/tylermeekel/egonote/internal/utils"
)

// noteRouter creates and returns a router for
// accessing the notes API
func noteRouter() *chi.Mux {
	r := chi.NewMux()

	r.Get("/{id}", getNote)

	return r
}

// getNote gets a note and returns it as JSON
// the way that DB is currently implemented makes
// this approach impossible
func getNote(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Write([]byte("Error"))
	}

	note := types.Note{
		ID:        id,
		Title:     "Hello",
		Content:   "This is content",
		Sharelink: "Linked",
	}
	utils.WriteJSON(w, note)
}

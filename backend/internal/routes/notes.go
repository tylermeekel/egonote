package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tylermeekel/egonote/internal/data"
	"github.com/tylermeekel/egonote/internal/types"
	"github.com/tylermeekel/egonote/internal/utils"
)

type NoteRouter struct {
	DB data.Database
}

// noteRouter creates and returns a router for
// accessing the notes API
func (n *NoteRouter) Routes() *chi.Mux {
	r := chi.NewMux()

	r.Post("/", n.postNote)
	r.Get("/{id}", n.getNote)
	r.Get("/", n.getNotes)
	r.Patch("/{id}", n.updateNote)
	r.Delete("/{id}", n.deleteNote)

	return r
}

func (n *NoteRouter) postNote(w http.ResponseWriter, r *http.Request) {
	var note types.Note

	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		utils.WriteJSON(w, "error") //! Temporary implementation
		return
	}

	noteResponse := n.DB.CreateNote(note)

	utils.WriteJSON(w, noteResponse)
}

// getNote gets a note and returns it as JSON
func (n *NoteRouter) getNote(w http.ResponseWriter, r *http.Request) {
	id := GetIDParam(r)
	if id < 0 {
		utils.WriteJSON(w, "error") //! Temporary implementation
		return
	}

	note := n.DB.GetNote(id)
	utils.WriteJSON(w, note) 
}

func (n *NoteRouter) getNotes(w http.ResponseWriter, r *http.Request) {
	notes := n.DB.GetNotes()
	utils.WriteJSON(w, notes)
}

func (n *NoteRouter) updateNote(w http.ResponseWriter, r *http.Request) {
	id := GetIDParam(r)
	if id < 0 {
		utils.WriteJSON(w, "error") //! Temporary implementation
		return
	}

	var note types.Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil{
		utils.WriteJSON(w, "error") //! Temporary implementation
		return
	}

	noteResponse := n.DB.UpdateNote(id, note)
	utils.WriteJSON(w, noteResponse)
}

func (n *NoteRouter) deleteNote(w http.ResponseWriter, r *http.Request) {
	id := GetIDParam(r)
	if id < 0 {
		utils.WriteJSON(w, "error") //! Temporary implementation
		return
	}

	note := n.DB.DeleteNote(id)
	utils.WriteJSON(w, note)
}

func (n *NoteRouter) createSharelink(w http.ResponseWriter, r *http.Request) {
	id := GetIDParam(r)
	if id < 0 {
		utils.WriteJSON(w, "error") //! Temporary implementation
		return
	}

	note := n.DB.GetNote(id)
	if note.Sharelink == ""{
		note.Sharelink = utils.CreateSharelink()
		n.DB.UpdateNote(id, note)
	}

	utils.WriteJSON(w, note)
}
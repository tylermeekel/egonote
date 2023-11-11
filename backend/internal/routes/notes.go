package routes

import (
	"encoding/json"
	"log"
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
		utils.WriteJSONError(w, "ID has to be a positive integer")
		return
	}

	noteResponse, err := n.DB.CreateNote(note)

	utils.WriteJSON(w, noteResponse)
}

// getNote gets a note and returns it as JSON
func (n *NoteRouter) getNote(w http.ResponseWriter, r *http.Request) {
	id := GetIDParam(r)
	if id < 0 {
		utils.WriteJSONError(w, "ID has to be a positive integer")
		return
	}

	note, err := n.DB.GetNote(id)
	if err != nil {
		log.Println(err.Error())
		utils.WriteInternalServerError(w)
		return
	}

	utils.WriteJSON(w, note)
}

func (n *NoteRouter) getNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := n.DB.GetNotes()
	if err != nil {
		log.Println(err.Error())
		utils.WriteInternalServerError(w)
		return
	}

	utils.WriteJSON(w, notes)
}

func (n *NoteRouter) updateNote(w http.ResponseWriter, r *http.Request) {
	id := GetIDParam(r)
	if id < 0 {
		utils.WriteJSONError(w, "ID has to be a positive integer")
		return
	}

	var note types.Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		log.Println(err.Error())
		utils.WriteInternalServerError(w)
		return
	}

	noteResponse, err := n.DB.UpdateNote(id, note)
	if err != nil {
		log.Println(err.Error())
		utils.WriteInternalServerError(w)
		return
	}

	utils.WriteJSON(w, noteResponse)
}

func (n *NoteRouter) deleteNote(w http.ResponseWriter, r *http.Request) {
	id := GetIDParam(r)
	if id < 0 {
		utils.WriteJSONError(w, "ID has to be a positive integer")
		return
	}

	note, err := n.DB.DeleteNote(id)
	if err != nil {
		log.Println(err.Error())
		utils.WriteInternalServerError(w)
		return
	}

	utils.WriteJSON(w, note)
}

func (n *NoteRouter) createSharelink(w http.ResponseWriter, r *http.Request) {
	id := GetIDParam(r)
	if id < 0 {
		utils.WriteJSONError(w, "ID has to be a positive integer")
		return
	}

	note, err := n.DB.GetNote(id)
	if err != nil {
		log.Println(err.Error())
		utils.WriteInternalServerError(w)
		return
	}

	if note.Sharelink == "" {
		note.Sharelink = utils.CreateSharelink()
		_, err = n.DB.UpdateNote(id, note)
		if err != nil {
			log.Println(err.Error())
			utils.WriteInternalServerError(w)
			return
		}
	}

	utils.WriteJSON(w, note)
}

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
	jw := utils.NewJSONResponseWriter(w)

	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		jw.AddError("id", "ID has to be a positive integer")
	}

	noteResponse, err := n.DB.CreateNote(note)
	if err != nil{
		jw.AddInternalError()
	}

	jw.AddData("note", noteResponse)
	jw.WriteJSON()
}

// getNote gets a note and returns it as JSON
func (n *NoteRouter) getNote(w http.ResponseWriter, r *http.Request) {
	jw := utils.NewJSONResponseWriter(w)
	id := GetIDParam(r)
	if id < 0 {
		jw.AddError("id", "ID must be a positive integer")
	}

	note, err := n.DB.GetNote(id)
	if err != nil {
		log.Println(err.Error())
		jw.AddInternalError()
	}

	jw.AddData("note", note)
	jw.WriteJSON()
}

func (n *NoteRouter) getNotes(w http.ResponseWriter, r *http.Request) {
	jw := utils.NewJSONResponseWriter(w)

	notes, err := n.DB.GetNotes()
	if err != nil {
		log.Println(err.Error())
		jw.AddInternalError()
	}

	jw.AddData("notes", notes)
	jw.WriteJSON()
}

func (n *NoteRouter) updateNote(w http.ResponseWriter, r *http.Request) {
	jw := utils.NewJSONResponseWriter(w)
	id := GetIDParam(r)
	if id < 0 {
		jw.AddError("id", "ID must be a positive integer")
	}

	var note types.Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		log.Println(err.Error())
		jw.AddInternalError()
	}

	noteResponse, err := n.DB.UpdateNote(id, note)
	if err != nil {
		log.Println(err.Error())
		jw.AddInternalError()
	}

	jw.AddData("note", noteResponse)
}

func (n *NoteRouter) deleteNote(w http.ResponseWriter, r *http.Request) {
	jw := utils.NewJSONResponseWriter(w)
	id := GetIDParam(r)
	if id < 0 {
		jw.AddError("id", "ID must be a positive integer")
	}

	note, err := n.DB.DeleteNote(id)
	if err != nil {
		log.Println(err.Error())
		jw.AddInternalError()
	}

	jw.AddData("note", note)
}

func (n *NoteRouter) createSharelink(w http.ResponseWriter, r *http.Request) {
	jw := utils.NewJSONResponseWriter(w)
	id := GetIDParam(r)
	if id < 0 {
		jw.AddError("id", "ID must be a positive integer")
	}

	note, err := n.DB.GetNote(id)
	if err != nil {
		log.Println(err.Error())
		jw.AddInternalError()
	}

	if note.Sharelink == "" {
		note.Sharelink = utils.CreateSharelink()
		_, err = n.DB.UpdateNote(id, note)
		if err != nil {
			log.Println(err.Error())
			jw.AddInternalError()
		}
	}

	jw.AddData("note", note)
	jw.WriteJSON()
}

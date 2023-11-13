package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tylermeekel/egonote/internal/auth"
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

	userID := auth.GetUserIDFromContext(r)
	if userID == -1{
		jw.AddError("user", "user not logged in")
		jw.WriteJSON()
		return
	}

	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		log.Println(err)
		jw.AddInternalError()
	}

	noteResponse, err := n.DB.CreateNote(userID, note)
	if err != nil{
		log.Println(err)
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

	userID := auth.GetUserIDFromContext(r)
	if userID == -1{
		jw.AddError("user", "user not logged in")
	}

	note, err := n.DB.GetNote(id, userID)
	if err != nil {
		log.Println(err.Error())
		jw.AddError("id", "Note not found")
	}

	jw.AddData("note", note)
	jw.WriteJSON()
}

func (n *NoteRouter) getNoteBySharelink(w http.ResponseWriter, r *http.Request) {
	jw := utils.NewJSONResponseWriter(w)
	sharelink := chi.URLParam(r, "sharelink")

	if sharelink == "" {
		jw.AddError("sharelink", "Must provide a sharelink")
		jw.WriteJSON()
		return
	}

	note, err := n.DB.GetNoteBySharelink(sharelink)
	if err != nil{
		log.Println(err)
		jw.AddError("sharelink", "Note not found")
	}

	jw.AddData("note", note)
	jw.WriteJSON()
}

func (n *NoteRouter) getNotes(w http.ResponseWriter, r *http.Request) {
	jw := utils.NewJSONResponseWriter(w)

	userID := auth.GetUserIDFromContext(r)
	if userID == -1{
		jw.AddError("user", "user not logged in")
	}

	notes, err := n.DB.GetNotes(userID)
	if err != nil {
		log.Println(err.Error())
		jw.AddInternalError()
	}

	jw.AddData("notes", notes)
	jw.WriteJSON()
}

func (n *NoteRouter) updateNote(w http.ResponseWriter, r *http.Request) {
	jw := utils.NewJSONResponseWriter(w)

	userID := auth.GetUserIDFromContext(r)
	if userID == -1{
		jw.AddError("user", "user not logged in")
	}

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

	noteResponse, err := n.DB.UpdateNote(id, userID, note)
	if err != nil {
		log.Println(err.Error())
		jw.AddInternalError()
	}

	jw.AddData("note", noteResponse)
	jw.WriteJSON()
}

func (n *NoteRouter) deleteNote(w http.ResponseWriter, r *http.Request) {
	jw := utils.NewJSONResponseWriter(w)

	userID := auth.GetUserIDFromContext(r)
	if userID == -1{
		jw.AddError("user", "user not logged in")
	}

	id := GetIDParam(r)
	if id < 0 {
		jw.AddError("id", "ID must be a positive integer")
	}

	note, err := n.DB.DeleteNote(id, userID)
	if err != nil {
		log.Println(err.Error())
		jw.AddInternalError()
	}

	jw.AddData("note", note)
	jw.WriteJSON()
}

func (n *NoteRouter) createSharelink(w http.ResponseWriter, r *http.Request) {
	jw := utils.NewJSONResponseWriter(w)

	userID := auth.GetUserIDFromContext(r)
	if userID == -1{
		jw.AddError("user", "user not logged in")
	}

	id := GetIDParam(r)
	if id < 0 {
		jw.AddError("id", "ID must be a positive integer")
	}

	note, err := n.DB.GetNote(id, userID)
	if err != nil {
		log.Println(err.Error())
		jw.AddInternalError()
	}

	if note.Sharelink == "" {
		note.Sharelink = utils.CreateSharelink()
		_, err = n.DB.UpdateNote(id, userID, note)
		if err != nil {
			log.Println(err.Error())
			jw.AddInternalError()
		}
	}

	jw.AddData("note", note)
	jw.WriteJSON()
}

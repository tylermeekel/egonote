package data

import "github.com/tylermeekel/egonote/internal/types"

// A Database is an interface for interacting
// with server data
type Database interface {
	CreateNote(types.Note) types.Note
	GetNotes() []types.Note
	GetNote(id int) types.Note
	UpdateNote(id int, note types.Note) types.Note
	DeleteNote(id int) types.Note
}

package data

import "github.com/tylermeekel/egonote/internal/types"

// A Database is an interface for interacting
// with server data
type Database interface {
	CreateNote(types.Note) (types.Note, error)
	GetNotes() ([]types.Note, error)
	GetNote(id int) (types.Note, error)
	UpdateNote(id int, note types.Note) (types.Note, error)
	DeleteNote(id int) (types.Note, error)

	CreateUser(username, password string) error
	DeleteUser(id int) error
	UpdateUser(id int, user types.User) error
}

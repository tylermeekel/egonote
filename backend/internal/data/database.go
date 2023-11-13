package data

import "github.com/tylermeekel/egonote/internal/types"

// A Database is an interface for interacting
// with server data
type Database interface {
	CreateNote(userID int, note types.Note) (types.Note, error)
	GetNotes(userID int) ([]types.Note, error)
	GetNote(id, userID int) (types.Note, error)
	GetNoteBySharelink(sharelink string) (types.Note, error)
	UpdateNote(id, userID int, note types.Note) (types.Note, error)
	DeleteNote(id, userID int) (types.Note, error)

	CreateUser(username, password string) (int, error)
	GetUser(username string) (types.User, error)
	DeleteUser(id int) error
	UpdateUser(id int, user types.User) error
}

package data

import "github.com/tylermeekel/egonote/internal/types"

// A Database is an interface for interacting
// with server data
type Database interface {
	GetNotes() []types.Note
	GetNote(id int) types.Note
}

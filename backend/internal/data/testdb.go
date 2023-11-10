package data

import (
	"math/rand"

	"github.com/tylermeekel/egonote/internal/types"
)

type TestDB struct{}

func (t *TestDB) CreateNote(note types.Note) types.Note{
	id := rand.Int()
	return types.Note{
		ID: id,
		Title: note.Title,
		Content: note.Content,
	}
}

func (t *TestDB) GetNote(id int) types.Note{
	return types.Note{
		ID: id,
		Title: "Title",
		Content: "Content",
		Sharelink: "123abc",
	}
} 

func (t *TestDB) GetNotes() []types.Note{
	var notes []types.Note
	for _, v := range []int{1,2,3}{
		note := types.Note{
			ID: v,
			Title: "Title",
			Content: "Content",
			Sharelink: "123abc",
		}
		notes = append(notes, note)
	}
	return notes
}

func (t *TestDB) UpdateNote(id int, note types.Note) types.Note{
	newNote := note
	newNote.ID = id
	return newNote
}

func (t *TestDB) DeleteNote(id int) types.Note{
	return types.Note{
		ID: id,
		Title: "Deleted Note",
		Content: "Deleted Content",
	}
}
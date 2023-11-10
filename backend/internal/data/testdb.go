package data

import "github.com/tylermeekel/egonote/internal/types"

type TestDB struct{}

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
package data

import (
	"database/sql"
	"log"
	"os"

	"github.com/tylermeekel/egonote/internal/types"
)

// Postgres is an implementation of a Database
// that has uses a Postgres connection for the
// server.
type Postgres struct {
	DB *sql.DB
}

// InitPostgres initializes a Postgres connection
// and returns a struct with the connection
func InitPostgres() *Postgres {
	uri := os.Getenv("POSTGRES_URI")
	db, err := sql.Open("postgres", uri)
	if err != nil {
		log.Fatalln("Error opening Postgres database: ", err.Error())
	}

	return &Postgres{
		DB: db,
	}
}

// GetNote gets a note from the Postgres server from a given ID
// and returns a types.Note
func (p *Postgres) GetNote(id int) types.Note {
	var note types.Note

	row := p.DB.QueryRow("SELECT * FROM notes WHERE id=$1", id)
	err := row.Scan(&note.ID, &note.Title, &note.Content, &note.Sharelink)
	if err != nil {
		log.Println(err.Error())
	}

	return note
}

// GetNotes gets all notes from the Postgres server and returns
// a slice of types.Note
func (p *Postgres) GetNotes() []types.Note {
	var notes []types.Note

	rows, err := p.DB.Query("SELECT * FROM notes")
	if err != nil {
		log.Println(err.Error())
	}
	for rows.Next() {
		var note types.Note
		rows.Scan(&note.ID, &note.Title, &note.Content, &note.Sharelink)
		notes = append(notes, note)
	}

	return notes
}

package data

import (
	"database/sql"
	"log"
	"os"

	"github.com/tylermeekel/egonote/internal/types"
	_ "github.com/lib/pq"
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

func (p *Postgres) CreateNote(note types.Note) (types.Note, error) {
	var createdNote types.Note
	
	row := p.DB.QueryRow("INSERT INTO notes(title, content) VALUES($1, $2) RETURNING *", note.Title, note.Content)
	err := row.Scan(&createdNote.ID, &createdNote.Title, &createdNote.Content, &createdNote.Sharelink)

	return createdNote, err
}

// GetNote gets a note from the Postgres server from a given ID
// and returns a types.Note
func (p *Postgres) GetNote(id int) (types.Note, error) {
	var note types.Note

	row := p.DB.QueryRow("SELECT * FROM notes WHERE id=$1", id)
	err := row.Scan(&note.ID, &note.Title, &note.Content, &note.Sharelink)

	return note, err
}

// GetNotes gets all notes from the Postgres server and returns
// a slice of types.Note
func (p *Postgres) GetNotes() ([]types.Note, error) {
	var notes []types.Note

	rows, err := p.DB.Query("SELECT * FROM notes")
	for rows.Next() {
		var note types.Note
		err = rows.Scan(&note.ID, &note.Title, &note.Content, &note.Sharelink)
		notes = append(notes, note)
	}

	return notes, err
}

func (p *Postgres) UpdateNote(id int, note types.Note) (types.Note, error) {
	var updatedNote types.Note
	
	query := 
	`UPDATE notes SET 
	title = COALESCE(NULLIF($1, ''), title),
	content = COALESCE(NULLIF($2, ''), content),
	sharelink = COALESCE(NULLIF($3, ''), sharelink)
	WHERE id=$4
	RETURNING *`
	
	row := p.DB.QueryRow(query, note.Title, note.Content, note.Sharelink, id)
	err := row.Scan(&updatedNote.ID, &updatedNote.Title, &updatedNote.Content, &updatedNote.Sharelink)

	return updatedNote, err
}

func (p * Postgres) DeleteNote(id int) (types.Note, error) {
	var deletedNote types.Note

	row := p.DB.QueryRow("DELETE FROM notes WHERE id=$1 RETURNING *", id)
	err := row.Scan(&deletedNote.ID, &deletedNote.Title, &deletedNote.Content, &deletedNote.Sharelink)

	return deletedNote, err
}

func (p *Postgres) CreateUser(username, password string) error {
	_, err := p.DB.Exec("INSERT INTO users(username, password) VALUES($1, $2)", username, password)

	return err
}

func (p *Postgres) GetUser(username string) (types.User, error) {
	var user types.User
	row := p.DB.QueryRow("SELECT * FROM users WHERE username=$1", username)
	err := row.Scan(&user.ID, &user.Username, &user.Password)

	return user, err
}

func (p *Postgres) UpdateUser(id int, user types.User) error {
	query :=
	`UPDATE notes SET 
	username = COALESCE(NULLIF($1, ''), username),
	password = COALESCE(NULLIF($2, ''), password),
	WHERE id=$3
	RETURNING *`

	_, err := p.DB.Exec(query, user.Username, user.Password, id)

	return err
}

func (p *Postgres) DeleteUser(id int) error {
	_, err := p.DB.Exec("DELETE FROM users WHERE id=$1", id)
	
	return err
}

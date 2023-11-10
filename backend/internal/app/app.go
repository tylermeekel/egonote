// Package app contains functions that help
// with the creation and running of the main
// server.
package app

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/tylermeekel/egonote/internal/data"
	"github.com/tylermeekel/egonote/internal/routes"
)

// An Application contains the routes and
// database connection for the server application
type Application struct {
	DB  data.Database
	Mux *chi.Mux
}

// CreateNewApp initializes a new *Application
// with a database and server connection and
// returns it.
func CreateNewApp(db data.Database) *Application {
	return &Application{
		DB:  db,
		Mux: routes.InitServerMux(),
	}
}

// StartApp starts the server
func StartApp() {
	godotenv.Load()

	port := os.Getenv("PORT")

	db := data.InitPostgres()
	app := CreateNewApp(db)

	log.Println("Server listening on", port)
	log.Fatalln(http.ListenAndServe(port, app.Mux))
}

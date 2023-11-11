// Package app contains functions that help
// with the creation and running of the main
// server.
package app

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/tylermeekel/egonote/internal/data"
	"github.com/tylermeekel/egonote/internal/routes"
)

// StartApp starts the server
func StartApp() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == ""{
		port = "3000"
	}

	db := data.InitPostgres()
	//db := &data.TestDB{}
	mux := routes.InitServerMux(db)

	log.Println("Server listening on", port)
	log.Fatalln(http.ListenAndServe(":"+port, mux))
}

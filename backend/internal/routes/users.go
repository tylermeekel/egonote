package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tylermeekel/egonote/internal/data"
)

type UserRouter struct {
	DB data.Database
}

func (u *UserRouter) Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/login", u.loginUser)
	r.Post("/register", u.registerUser)

	return r
}

func (u *UserRouter) loginUser(w http.ResponseWriter, r *http.Request) {

}

func (u *UserRouter) registerUser(w http.ResponseWriter, r *http.Request) {

}

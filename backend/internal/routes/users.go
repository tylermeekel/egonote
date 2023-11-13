package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/tylermeekel/egonote/internal/auth"
	"github.com/tylermeekel/egonote/internal/data"
	"github.com/tylermeekel/egonote/internal/types"
	"github.com/tylermeekel/egonote/internal/utils"
	"golang.org/x/crypto/bcrypt"
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
	jw := utils.NewJSONResponseWriter(w)
	var givenCredentials types.User
	err := json.NewDecoder(r.Body).Decode(&givenCredentials)
	if err != nil {
		log.Println(err.Error())
		jw.AddInternalError()
	}

	dbCredentials, err := u.DB.GetUser(givenCredentials.Username)
	if err != nil {
		log.Println(err.Error())
		jw.AddError("login", "Username or Password Incorrect")
		jw.WriteJSON()
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbCredentials.Password), []byte(givenCredentials.Password))
	if err != nil{
		log.Println(err.Error())
		jw.AddError("login", "Username or Password Incorrect")
		jw.WriteJSON()
		return
	}


	expTime := time.Now().Add(5*time.Hour)
	signedString, err := auth.SignJWT(dbCredentials.Username, expTime)

	jwtCookie := http.Cookie{
		Name: "jwt_token",
		Value: signedString,
		Path: "/",
		Expires: expTime,
	}
	http.SetCookie(w, &jwtCookie)
	jw.AddData("login", "login successful")
}

func (u *UserRouter) registerUser(w http.ResponseWriter, r *http.Request) {
	jw := utils.NewJSONResponseWriter(w)
	var givenCredentials types.User
	json.NewDecoder(r.Body).Decode(&givenCredentials)

	givenCredentials.Username = strings.ToLower(givenCredentials.Username)

	if !types.ValidateUser(givenCredentials) {
		jw.AddError("logininfo", "Username or Password doesn't meet guidelines") //! Add new implementation of WriteJSONError for users specifically
		jw.WriteJSON()
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(givenCredentials.Password), 10)
	if err != nil{
		log.Println(err.Error())
		jw.AddInternalError()
		jw.WriteJSON()
		return
	}

	err = u.DB.CreateUser(givenCredentials.Username, string(hashedPassword))
	if err != nil{
		log.Println(err.Error())
		jw.AddInternalError()
		jw.WriteJSON()
		return
	}
	
	expTime := time.Now().Add(5*time.Hour)
	signedString, err := auth.SignJWT(givenCredentials.Username, expTime)
	if err != nil{
		log.Println(err.Error())
		jw.AddInternalError()
		jw.WriteJSON()
		return
	}

	jwtCookie := http.Cookie{
		Name: "jwt_token",
		Value: signedString,
		Path: "/",
		Expires: expTime,
	}
	http.SetCookie(w, &jwtCookie)
	jw.AddData("register", "register successful")
}

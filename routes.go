package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log"
	"bytes"
	"net/http"
)

func router() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.Post("/register", handleRegister)
	router.Post("/login", handleLogin)

	return router
}

func handleRegister(res http.ResponseWriter, req *http.Request) {
	var (
		user User
		err  error
	)

	if err = render.DecodeJSON(req.Body, &user); err != nil {
		log.Println("Failed to decode json with error:", err)
		res.WriteHeader(400)
		return
	}

	if user.Salt, user.PassHash, err = SaltAndHash(user.Password); err != nil {
		log.Println("Failed to salt and hash password:", err)
		res.WriteHeader(500)
		return
	}

	user.Password = ""

	if _, err = db.InsertUser(user); err != nil {
		log.Println("Failed to add user to database:", err)
		res.WriteHeader(500)
		return
	}

	res.WriteHeader(200)
}

func handleLogin(res http.ResponseWriter, req *http.Request) {
	var (
		response LoginResponse
		auth bool
		username string
		password string
		user User
		hash []byte
		err error
	)

	if username, password, auth = req.BasicAuth(); !auth {
		log.Println("No authorization provided")
		res.WriteHeader(400)
		return
	}

	user, err = db.GetUserByName(username)
	if err != nil {
		log.Println("User", username, "not found.", err.Error())
		res.WriteHeader(401)
		return
	}
		
	hash, err = HashPassword(password, user.Salt)
	if err != nil {
		log.Println("Failed to hash password")
		res.WriteHeader(500)
		return
	}
	
	if bytes.Compare(hash, user.PassHash) != 0 {
		log.Println("Invalid login")
		res.WriteHeader(401)
		return
	}
	
	response.Token, err = NewToken(user)
	if err != nil {
		log.Println("Failed to generate token")
		res.WriteHeader(500)
		return
	}

	render.JSON(res, req, response)
}

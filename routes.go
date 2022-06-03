package main

import (
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func router() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.Post("/register", handleRegister)
	
	return router
}

func handleRegister(res http.ResponseWriter, req *http.Request) {
	var (
		user User
		err error
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


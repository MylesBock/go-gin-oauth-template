package main

import (
	"callback"
	"home"
	"log"
	"login"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	r := mux.NewRouter()

	r.HandleFunc("/", home.HomeHandler)
	r.HandleFunc("/callback", callback.CallbackHandler)
	r.HandleFunc("/login", login.LoginHandler)
	http.Handle("/", r)

	log.Print("Server listening on http://localhost:3000/")
	log.Fatal(http.ListenAndServe("0.0.0.0:3000", nil))

}

/*
r.HandleFunc("/login", login.LoginHandler)
		r.HandleFunc("/logout", logout.LogoutHandler)
		r.Handle("/user", negroni.New(
		negroni.HandlerFunc(middlewares.IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(user.UserHandler)),
	))

		r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))

*/

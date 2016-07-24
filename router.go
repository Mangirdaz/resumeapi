package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func NewRouter() {

	//create new router
	router := mux.NewRouter().StrictSlash(false)
	storage := NewLibKVBackend()

	//api backend init
	apiV1 := router.PathPrefix("/api/v1").Subrouter()

	apiV1.Methods("GET").Path("").Name("Index").Handler(http.HandlerFunc(mybackendHandler(Index, storage)))

	//resume methods
	resume := apiV1.PathPrefix("/resume").Subrouter()
	resume.Methods("GET").Name("Resume").Handler(http.HandlerFunc(mybackendHandler(Resume, storage)))
	resume.Methods("GET").Path("/{1}").Name("Resume").Handler(http.HandlerFunc(mybackendHandler(Resume, storage)))

	notes := apiV1.PathPrefix("/notes").Subrouter()
	notes.Methods("GET").Path("/{key}").Name("Notes").Handler(http.HandlerFunc(mybackendHandler(GetNote, storage)))
	notes.Methods("GET").Name("Notes").Handler(http.HandlerFunc(mybackendHandler(GetNotes, storage)))
	notes.Methods("POST").Name("AddNotes").Handler(http.HandlerFunc(mybackendHandler(AddNotes, storage)))

	//middleware intercept
	midd := http.NewServeMux()
	midd.Handle("/", router)
	midd.Handle("/api/v1/notes", negroni.New(
		negroni.HandlerFunc(CheckAuth),
		negroni.Wrap(apiV1),
	))
	n := negroni.Classic()
	n.UseHandler(midd)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), n))

	//return router

}

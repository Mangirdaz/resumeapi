package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {

	//create new router
	router := mux.NewRouter().StrictSlash(true)
	storage := NewLibKVBackend()

	//api backend init
	api := router.PathPrefix("/api/v1").Subrouter()

	//resume methods
	api.Methods("GET").Path("/").Name("Index").Handler(http.HandlerFunc(mybackendHandler(Index, storage)))
	api.Methods("GET").Path("/resume").Name("Resume").Handler(http.HandlerFunc(mybackendHandler(Resume, storage)))
	api.Methods("GET").Path("/resume/{1}").Name("Resume").Handler(http.HandlerFunc(mybackendHandler(Resume, storage)))

	//notes methods
	api.Methods("GET").Path("/notes/{key}").Name("Notes").Handler(http.HandlerFunc(mybackendHandler(GetNote, storage)))
	api.Methods("GET").Path("/notes").Name("Notes").Handler(http.HandlerFunc(mybackendHandler(GetNotes, storage)))

	return router
}

package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"1",
		"GET",
		"/resume/{1}",
		TodoIndex,
	},
	Route{
		"Notes",
		"GET",
		"/notes",
		GetNotes,
	},
	Route{
		"Notes",
		"GET",
		"/notes/{key}",
		GetNote,
	},
	Route{
		"Notes",
		"POST",
		"/notes",
		AddNote,
	},
}

package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// change to /{section}/{subsection}
// Route{
// 		"TodoCreate",
// 		"GET",
// 		"/volunteer",
// 		TodoCreate,
// 	},
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
		"/{1}/{2}/{3}",
		TodoIndex,
	},
	Route{
		"1",
		"GET",
		"/{1}/{2}",
		TodoIndex,
	},
	Route{
		"1",
		"GET",
		"/{1}",
		TodoIndex,
	},
}

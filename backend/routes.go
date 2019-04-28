package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	Query       bool
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"ListAnnouncements",
		"GET",
		"/",
		false,
		listAnnouncements,
	},
	Route{
		"CreateAnnouncement",
		"POST",
		"",
		true,
		createAnnouncement,
	},
}

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		if route.Query == true {
			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Queries("start", "{start:[0-9]+}", "end", "{end:[0-9]+}").
				Handler(handler)
		} else {
			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(handler)
		}
	}
	return router
}

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf("%s\t%s\t%s\t%s", r.Method, r.RequestURI, name, time.Since(start))
	})
}

package main

import "net/http"

func listAnnouncements(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Esterit!\n"))
}

func createAnnouncement(w http.ResponseWriter, r *http.Request) {

}

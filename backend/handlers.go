package main

import "net/http"
import "fmt"

func listAnnouncements(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Go Esteri, Go!")

}

func createAnnouncement(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Go Esteri, Go!")
}

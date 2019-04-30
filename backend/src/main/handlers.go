package main

import (
	"html/template"
	"net/http"
)

func listAnnouncements(w http.ResponseWriter, r *http.Request) {
	announcements := selectAnnouncements()

	for _, announcement := range announcements {
		if announcement.Announcement != "" {
			temppi := template.New("jotain")
			temppi, _ = temppi.Parse("Category: {{ .Category }} \nName: {{ .Username }} \nEmail: {{ .Useremail }}\nAnnouncement: {{ .Announcement }}\nExpiration date: {{ .ExpirationDate }}\n- - -\n\n")

			temppi.Execute(w, announcement)
		}
	}

}

func createAnnouncement(w http.ResponseWriter, r *http.Request) {
	var newAnnouncement Announcement
	newAnnouncement.Announcement = r.FormValue("description")
	newAnnouncement.Category = r.FormValue("category")
	newAnnouncement.ExpirationDate = r.FormValue("date")
	newAnnouncement.Useremail = r.FormValue("email")
	newAnnouncement.Username = r.FormValue("name")
	insertAnnouncement(newAnnouncement)
}

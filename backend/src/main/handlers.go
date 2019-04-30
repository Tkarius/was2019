package main

import "net/http"
import "fmt"
import "html/template"

func listAnnouncements(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Go Esteri, Go!")
	announcements := selectAnnouncements()

	for _, announcement := range announcements {
		temppi := template.New("jotain")
		temppi, _ = temppi.Parse("{{.}}<div class='bnsannouncement'><div class='firstline'><span class='title'>{{ .Title }}</span> - <span class='date'> {{ .Date }} </span> - {{ .Useremail }} - {{ .Username }}</div><span class='category'>{{ .Category }}</span><div class='description'>{{ .Announcement }}</div></div>")
		/*fmt.Println("Name")
		fmt.Println(announcement.Username)
		fmt.Println("Email")
		fmt.Println(announcement.Useremail)
		fmt.Println("Category:")
		fmt.Println(announcement.Category)
		fmt.Println("Announcement")
		fmt.Println(announcement.Announcement)
		fmt.Println("Expiration date:")
		fmt.Println(announcement.ExpirationDate)*/
		temppi.Execute(w, announcement)
	}

}

func createAnnouncement(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Go Esteri, Go!")
}

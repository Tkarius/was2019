package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Announcement struct {
	Category       string
	Announcement   string
	Username       string
	Useremail      string
	ExpirationDate string
	Title          string
}

// global :s Let's see if we need to leave it that way.
// this is not exposed outside database.go.
var db *sql.DB

func selectAnnouncements() []Announcement {
	fmt.Printf("DEBUG: Connecting to db with view user.")

	var err error
	db, err = sql.Open("sqlserver", env.selectUserSecret)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		fmt.Println("DEBUG: Error with db connection")
		fmt.Println(err)
	}

	rows, err := db.QueryContext(ctx, "was_groupwork.SelectBusinessAnnouncements")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var name, email, announcement, category string
	var expirationDate string
	announcements := make([]Announcement, 20)
	var announceIndex int
	announceIndex = 0

	for rows.Next() {
		if err := rows.Scan(&name, &email, &announcement, &category, &expirationDate); err != nil {
			fmt.Println("Something bad happened while scanning db results.")
			fmt.Println(err)
		}
		// VALIDATE STUFFI!
		fmt.Printf("name: %s email: %s announcement: %s category: %s expiration date: %s\n", name, email, announcement, category, expirationDate)
		announcements[announceIndex].Category = category
		announcements[announceIndex].Announcement = announcement
		announcements[announceIndex].Username = name
		announcements[announceIndex].Useremail = email
		announcements[announceIndex].ExpirationDate = expirationDate
		announceIndex++
	}
	return announcements
}

func insertAnnouncement() {
	// VALIDATE STUFFI!
	var err error
	db, err = sql.Open("sqlserver", env.insertUserSecret)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		fmt.Println("DEBUG: Error with db connection")
		fmt.Println(err)
	}

	timestamp := time.Now()
	timestamp.Add(time.Hour * 3)

	_, err = db.ExecContext(ctx, "was_groupwork.CreateBusinessAnnouncement",
		sql.Named("category", "buying"),
		sql.Named("announcement", "haluisin kovasti ostaa jotain"),
		sql.Named("username", "esterihanse"),
		sql.Named("useremail", "esteri@testeri.fi"),
		sql.Named("expiration_date", timestamp),
	)
	if err != nil {
		fmt.Println(err)
	}

}

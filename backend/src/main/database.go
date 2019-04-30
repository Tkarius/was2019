package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
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
	var err error
	db, err = sql.Open("sqlserver", env.selectUserSecret)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		panic(err)
	}

	rows, err := db.QueryContext(ctx, "was_groupwork.SelectBusinessAnnouncements")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var name, email, announcement, category string
	var expirationDate string
	announcements := make([]Announcement, 50)
	var announceIndex int
	announceIndex = 0

	for rows.Next() {
		if err := rows.Scan(&name, &email, &announcement, &category, &expirationDate); err != nil {
			fmt.Println("Something bad happened while scanning db results.")
			fmt.Println(err)
		}
		// VALIDATE STUFFI!
		announcements[announceIndex].Category = category
		announcements[announceIndex].Announcement = announcement
		announcements[announceIndex].Username = name
		announcements[announceIndex].Useremail = email
		announcements[announceIndex].ExpirationDate = expirationDate
		announceIndex++
	}
	return announcements
}

func insertAnnouncement(toInsert Announcement) {
	// VALIDATE STUFFI!
	var err error
	db, err = sql.Open("sqlserver", env.insertUserSecret)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		panic(err)
	}

	parsedTime, _ := strconv.ParseInt(toInsert.ExpirationDate, 10, 64)
	timeStamp := time.Unix(parsedTime/1000, 0)

	_, err = db.ExecContext(ctx, "was_groupwork.CreateBusinessAnnouncement",
		sql.Named("category", toInsert.Category),
		sql.Named("announcement", toInsert.Announcement),
		sql.Named("username", toInsert.Username),
		sql.Named("useremail", toInsert.Useremail),
		sql.Named("expiration_date", timeStamp),
	)
	if err != nil {
		fmt.Println(err)
	}
}

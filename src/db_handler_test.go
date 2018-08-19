package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

type DBTuple struct {
	id         int
	username   string
	url        string
	hashedName uint32
}

func TestSelectAll(t *testing.T) {
	db, err := sql.Open("sqlite3", "./instaGo2Channel.db")
	checkErr(err)

	query, err := db.Query("SELECT * FROM info_test WHERE username=\"user1\"")
	checkErr(err)
	var id int
	var username string
	var url string
	var hashedName int

	dbTuple := []DBTuple{
		DBTuple{1, "user1", "inst.com/1.jpg", 1234},
		DBTuple{2, "user1", "inst.com/2.jpg", 5678},
		DBTuple{3, "user1", "inst.com/1.mp4", 12345678},
	}

	counter := 0
	for query.Next() {
		err = query.Scan(&id, &username, &url, &hashedName)
		checkErr(err)

		idEquals := id == dbTuple[counter].id
		usernameEquals := username == dbTuple[counter].username
		urlEquals := url == dbTuple[counter].url
		hashedNameEquals := uint32(hashedName) == dbTuple[counter].hashedName

		counter++
		if !(idEquals && usernameEquals && urlEquals && hashedNameEquals) {
			t.Error("\"Select * ... \" was fetch wrong data")
		}
	}
	query.Close()
}

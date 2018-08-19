package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type DBTuple struct {
	username   string
	url        string
	hashedName int
}

func pasteIntoDB(db *sql.DB, stories *Stories){
	fetchedTuples := fetchAllByUserName(db, stories.username)
	hashedNames := hashingNames(stories.urls)

	counter := 0
	for i, hashedName := range hashedNames {
		counter = 0
		for _, tuple := range *fetchedTuples {
			if int(hashedName) != tuple.hashedName {
				counter++
			}
		}
		if counter == len(*fetchedTuples) {
			inst, err := db.Prepare("INSERT INTO info(username, url, hashed_name) values(?,?,?)")
			checkErr(err)
			_, err = inst.Exec(stories.username, stories.urls[i], hashedName)
			checkErr(err)
		}
	}
}

func fetchAllByUserName(db *sql.DB, userName string) *[]DBTuple {
	var fetchedTuples []DBTuple

	queryString := fmt.Sprintf("SELECT * FROM info WHERE username=\"%s\"", userName)
	query, err := db.Query(queryString)
	checkErr(err)
	var id int
	var username string
	var url string
	var hashedName int

	for query.Next() {
		err = query.Scan(&id, &username, &url, &hashedName)
		checkErr(err)
		fetchedTuples = append(fetchedTuples, DBTuple{username, url, hashedName})
	}
	query.Close()
	return &fetchedTuples;
}
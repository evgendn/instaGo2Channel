package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type DBTuple struct {
	username   string
	url        string
	hashedName uint32
}

func storiesNotExist(db *sql.DB, dbTuple *DBTuple) bool {
	fetchedTuples := fetchAllByUserName(db, dbTuple.username)
	hashedName := hashingName(dbTuple.url)
	dbTuple.hashedName = hashedName

	hits := 0
	for _, tuple := range *fetchedTuples {
		if hashedName != tuple.hashedName {
			hits++
		}
	}

	if hits == len(*fetchedTuples) {
		return true
	}
	return false
}

func pasteIntoDB(db *sql.DB, dbTuple *DBTuple) {
	inst, err := db.Prepare("INSERT INTO info(username, url, hashed_name) values(?,?,?)")
	checkErr(err)

	username := dbTuple.username
	url := dbTuple.url
	hashedName := dbTuple.hashedName

	_, err = inst.Exec(username, url, hashedName)
	checkErr(err)
}

func fetchAllByUserName(db *sql.DB, userName string) *[]DBTuple {
	var fetchedTuples []DBTuple

	queryString := fmt.Sprintf("SELECT username, url, hashed_name FROM info WHERE username=\"%s\"", userName)
	query, err := db.Query(queryString)
	checkErr(err)

	var username string
	var url string
	var hashedName uint32

	for query.Next() {
		err = query.Scan(&username, &url, &hashedName)
		checkErr(err)
		fetchedTuples = append(fetchedTuples, DBTuple{username, url, hashedName})
	}
	query.Close()
	return &fetchedTuples;
}

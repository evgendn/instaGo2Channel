package main

import (
	"os"
	"database/sql"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	username := os.Getenv("NICKNAME")
	if username == "" {
		panic("$NICKNAME does not exist")
	}

	stories := getStories(username)
	storiesUrls := stories.urls

	db, err := sql.Open("sqlite3", "./instaGo2Channel.db")
	checkErr(err)

	dbTuple := DBTuple{
		username: stories.username,
		url: "",
		hashedName: 0,
	}
	for _, url := range storiesUrls {
		dbTuple.url = url
		if storiesNotExist(db, &dbTuple) {
			pasteIntoDB(db, &dbTuple)
			downloadMedia(dbTuple.url)
			filename := getNameFromURL(dbTuple.url)
			if _, err := os.Stat(filename); err == nil {
				bot, err := tgbotapi.NewBotAPI(os.Getenv("ACCESS_TOKEN"))
				checkErr(err)

				fileFormat := getFileFormat(filename)
				if fileFormat == "jpg" {
					sendImage(bot, filename, "#stories")
				} else if fileFormat == "mp4" {
					sendVideo(bot, filename, "#stories")
				}
				
				if err := os.Remove(filename); err != nil {
       					panic(err)
    				}
			}
		}
	}
}

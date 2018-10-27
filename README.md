# InstaGo2Channel

This project allows crossposting stories from instagram to telegram channel, including photos and videos.

*Created for learning golang.*

## How to

1. Install packages
```sh
go get github.com/mattn/go-sqlite3
go get -u github.com/go-telegram-bot-api/telegram-bot-api
```

2. Write crontab config, this one is running every 12 hours 
```sh
NICKNAME="Instagram username"
ACCESS_TOKEN="Telegram bots access token"
CHANNEL_NAME="@TelegramChannelName"
* */12 * * * cd ~/Path/to/bin/file && ./Go_Run
```
3. Create sqlite db and name it "instaGo2Channel.db"
```sql
CREATE TABLE "info" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    "username" TEXT NOT NULL,
    "url" TEXT NOT NULL,
    "hashed_name" INTEGER NOT NULL
)
```
4. For database tests you have to create the same table, but name it "info_test"
```sql
CREATE TABLE "info_test" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    "username" TEXT NOT NULL,
    "url" TEXT NOT NULL,
    "hashed_name" INTEGER NOT NULL
)

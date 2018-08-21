# InstaGo2Channel v0.1

This project allows crossposting stories from instagram to telegram channel, including photos and videos.

*Created for learning golang.*

## How to

1. Set crontab config, every 12 hours 
```sh
NICKNAME="Instagram username"
ACCESS_TOKEN="Telegram bots access token"
CHANNEL_NAME="@TelegramChannelName"
* */12 * * * cd ~/Path/to/bin/file && ./Go_Run
```
2. Create sqlite db and name it "instaGo2Channel.db"
```sql
CREATE TABLE "info" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    "username" TEXT NOT NULL,
    "url" TEXT NOT NULL,
    "hashed_name" INTEGER NOT NULL
)
```
3. For database tests you have to create the same table, but name it "info_test"

## To Do List
- [x] Add cron task
- [ ] Remove media after sending to telegram 
- [ ] Add goroutines

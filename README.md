# InstaGo2Channel

This project allows crossposting stories from instagram to telegram channel, including photos and videos.

*Created for learning golang.*

## How to

1. Set environment variables
```sh
export NICKNAME="Instagram username"
export ACCESS_TOKEN="Telegram bot access token"
export CHANNEL_NAME="Telegram channel name"
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

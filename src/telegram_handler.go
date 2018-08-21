package main

import (
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func sendImage(bot *tgbotapi.BotAPI, pathToImage, caption string) {
	msg := tgbotapi.PhotoConfig{
		BaseFile: tgbotapi.BaseFile{
			BaseChat: tgbotapi.BaseChat{ChannelUsername: os.Getenv("CHANNEL_NAME")},
			File: pathToImage,
		},
		Caption: caption,
	}
	bot.Send(msg)
}

func sendVideo(bot *tgbotapi.BotAPI, pathToVideo, caption string) {
	msg := tgbotapi.VideoConfig{
		BaseFile: tgbotapi.BaseFile{
			BaseChat: tgbotapi.BaseChat{ChannelUsername: os.Getenv("CHANNEL_NAME")},
			File: pathToVideo,
		},
		Caption: caption,
	}
	bot.Send(msg)
}

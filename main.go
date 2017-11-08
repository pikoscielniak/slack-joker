package main

import (
	"os"
	"slack-joker/bot"
)

func main() {
	token := os.Getenv("SLACK_TOKEN")
	bot.Start(token)
}

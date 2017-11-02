package main

import (
	"slack-joker/fetcher"
	"slack-joker/sender"
	"os"
)

func main() {
	joke, err := fetcher.Fetch()
	if err != nil {
		panic(err)
	}
	channel := os.Args[1]
	token := os.Args[2]
	err = sender.SendToSlack(joke, channel, token)
	if err != nil {
		panic(err)
	}
}

package joker

import (
	"github.com/nlopes/slack"
	"log"
)

const (
	maxJokesToRemove = 10
)

var sendJokesTimeStamps = newRotateQueue(maxJokesToRemove)

func JokeCommand(ev *slack.MessageEvent, rtm *slack.RTM) {
	log.Println("Handler defaultJokeCommand")
	joke, err := fetch()
	if err != nil {
		log.Println(err)
	}
	params := slack.PostMessageParameters{
		Username: rtm.GetInfo().User.Name,
	}
	_, respTimeStamp, _ := rtm.PostMessage(ev.Channel, joke, params)
	sendJokesTimeStamps.add(respTimeStamp)
}

package joker

import (
	"github.com/nlopes/slack"
	"log"
)

func SucharCommand(ev *slack.MessageEvent, rtm *slack.RTM) {
	log.Println("Handler sucharCommand")
	respTimeStamp := sendJokesTimeStamps.pop()
	if respTimeStamp != "" {
		rtm.DeleteMessage(ev.Channel, respTimeStamp)
	} else {
		rtm.SendMessage(rtm.NewOutgoingMessage("nie wiem co masz na myśli", ev.Channel))
	}
}

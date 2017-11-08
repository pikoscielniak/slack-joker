package bot

import (
	"github.com/nlopes/slack"
	"fmt"
	"strings"
	"log"
	"slack-joker/botAbstract"
	"slack-joker/joker"
)

func Start(token string) {
	api := slack.New(token)
	rtm := api.NewRTM()
	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			log.Println("Event Received: " + msg.Type)
			switch ev := msg.Data.(type) {

			case *slack.ConnectedEvent:
				log.Println("Connection counter:", ev.ConnectionCount)

			case *slack.MessageEvent:
				handleMessageEvent(ev, rtm)

			case *slack.RTMError:
				log.Println("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				log.Println("Invalid credentials")
				break Loop

			default:
				//Take no action
			}
		}
	}
}

var commandHandlers = map[string]botAbstract.Command{
	"suchar": joker.SucharCommand,
}

func greetingCommand(ev *slack.MessageEvent, rtm *slack.RTM) {
	log.Println("Handler greetingCommand")
	rtm.SendMessage(rtm.NewOutgoingMessage("hej", ev.Channel))
}

func handleMessageEvent(ev *slack.MessageEvent, rtm *slack.RTM) {
	log.Printf("Message: %v\n", ev)
	info := rtm.GetInfo()
	log.Printf("User: %s\n", info.User.Name)
	prefix := fmt.Sprintf("<@%s> ", info.User.ID)

	if ev.User != info.User.ID && strings.HasPrefix(ev.Text, prefix) {
		parts := strings.Fields(ev.Text)
		if len(parts) == 1 {
			greetingCommand(ev, rtm)
		}
		cmd := parts[1]
		cmdHand, ok := commandHandlers[cmd]
		if !ok {
			cmdHand = joker.JokeCommand
		}
		cmdHand(ev, rtm)
	}
}

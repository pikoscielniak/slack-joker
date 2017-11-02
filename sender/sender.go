package sender

import (
	"github.com/nlopes/slack"
)

func SendToSlack(msg, channel, token string) error {
	api := slack.New(token)
	params := slack.PostMessageParameters{
		Username: "Joker",
	}
	_, _, err := api.PostMessage(channel, msg, params)
	if err != nil {
		return err
	}
	return nil
}

package botAbstract

import "github.com/nlopes/slack"

type Command func(ev *slack.MessageEvent, rtm *slack.RTM)

package plugins

import (
	"github.com/nlopes/slack"
)

type Plugin interface {
	Match(string, slack.PostMessageParameters) slack.PostMessageParameters
	GetName() string
}

type Match func(string, slack.PostMessageParameters) slack.PostMessageParameters

var Registry = []Plugin{}

func Add(p Plugin) {
	Registry = append(Registry, p)
}

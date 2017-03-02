package ping

import (
	"github.com/domudall/doiici/plugins"
	"github.com/nlopes/slack"
)

type plugin struct{}

func (p *plugin) Match(input string, params slack.PostMessageParameters) slack.PostMessageParameters {
	return slack.PostMessageParameters{
		Text: "pong!",
	}
}

func (p *plugin) GetName() string {
	return "ping"
}

func init() {
	plugins.Add(&plugin{})
}

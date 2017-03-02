package echo

import (
	"fmt"

	"github.com/domudall/doiici/plugins"
	"github.com/nlopes/slack"
)

type plugin struct{}

func (p *plugin) Match(input string, params slack.PostMessageParameters) slack.PostMessageParameters {
	if input == "" {
		return p.Help(params)
	}

	return slack.PostMessageParameters{
		Text: fmt.Sprintf("echo: %s", input),
	}
}

func (p *plugin) GetName() string {
	return "echo"
}

func (p *plugin) Help(params slack.PostMessageParameters) slack.PostMessageParameters {
	return slack.PostMessageParameters{
		Text: "pong!",
	}
}

func init() {
	plugins.Add(&plugin{})
}

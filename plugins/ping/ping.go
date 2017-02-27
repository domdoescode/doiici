package ping

import (
	"github.com/domudall/doiici/plugins"
)

type plugin struct{}

func (p *plugin) Match(input string) string {
	return "pong!"
}

func (p *plugin) GetName() string {
	return "ping"
}

func init() {
	plugins.Add(&plugin{})
}

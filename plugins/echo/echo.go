package echo

import (
	"fmt"

	"github.com/domudall/doiici/plugins"
)

type plugin struct{}

func (p *plugin) Match(input string) string {
	if input == "" {
		return p.Help()
	}

	return fmt.Sprintf("echo: %s", input)
}

func (p *plugin) GetName() string {
	return "echo"
}

func (p *plugin) Help() string {
	return "TODO"
}

func init() {
	plugins.Add(&plugin{})
}

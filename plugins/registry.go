package plugins

type Plugin interface {
	Match(string) string
	GetName() string
}

type Match func(string) string

var Registry = []Plugin{}

func Add(p Plugin) {
	Registry = append(Registry, p)
}

package action

type Action interface {
	Callback() func()
	Command() string
	ResponseParser(string) error
}

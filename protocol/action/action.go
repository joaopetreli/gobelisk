package action

type Action interface {
	Query() string
	Parse(response string)
	Callback()
}

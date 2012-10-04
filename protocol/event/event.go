package event

type Event interface {
	Parse(string)
	Callback()
}

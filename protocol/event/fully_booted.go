package event

import (
	"fmt"
	"gobelisk/protocol"
	"strings"
)

type FullyBooted struct {
	Event       string
	Privilege   string
	Status      string
	RawResponse string
	callback    func(FullyBooted)
}

func NewFullyBooted() FullyBooted {
	var fullyBooted FullyBooted
	fullyBooted.Event = "FullyBooted"
	fullyBooted.callback = func(fb FullyBooted) {
		fmt.Print(fb.RawResponse)
	}

	return fullyBooted
}

func (fb *FullyBooted) Parse(response string) error {
	fb.RawResponse = response
	if len(response) != 67 {
		return protocol.ErrInvalidResponse
	}

	lines := strings.Split(response, "\r\n")

	fb.Event = lines[0][7:]
	fb.Privilege = lines[1][11:]
	fb.Status = lines[2][8:]

	return nil
}

func (fb FullyBooted) Callback() {
	fb.callback(fb)
}

func (fb *FullyBooted) SetCallback(f func(FullyBooted)) {
	fb.callback = f
}

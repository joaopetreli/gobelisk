package action

import (
	"fmt"
	"gobelisk/protocol"
	"strings"
)

type Logoff struct {
	Action   string
	callback func(Logoff)
	LogoffResponse
}

type LogoffResponse struct {
	Response    string
	Message     string
	RawResponse string
}

func NewLogoff() Logoff {
	var logoff Logoff

	logoff.Action = "Logoff"
	logoff.callback = func(l Logoff) {
		fmt.Print(l.RawResponse)
	}

	return logoff
}

func (l Logoff) Query() string {
	return fmt.Sprintf("Action: %s\r\n\r\n", l.Action)
}

func (l *Logoff) Parse(response string) error {
	l.RawResponse = response

	if len(response) != 56 {
		return protocol.ErrInvalidResponse
	}

	lines := strings.Split(response, "\r\n")

	l.Response = lines[0][10:]
	l.Message = lines[1][9:]

	return nil
}

func (l Logoff) Callback() {
	l.callback(l)
}

func (l *Logoff) SetCallback(f func(logoff Logoff)) {
	l.callback = f
}

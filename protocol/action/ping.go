package action

import (
	"fmt"
	"gobelisk/protocol"
	"strings"
)

type Ping struct {
	Action   string
	callback func(Ping)
	PingResponse
}

type PingResponse struct {
	Response    string
	Ping        string
	Timestamp   string
	RawResponse string
}

func NewPing() Ping {
	var ping Ping

	ping.Action = "Ping"
	ping.callback = func(p Ping) {
		fmt.Print(p.RawResponse)
	}

	return ping
}

func (p Ping) Query() string {
	return fmt.Sprintf("Action: %s\r\n\r\n", p.Action)
}

func (p *Ping) Parse(response string) error {
	p.RawResponse = response

	if len(response) != 63 {
		return protocol.ErrInvalidResponse
	}

	lines := strings.Split(response, "\r\n")
	p.Response = lines[0][10:]
	p.Ping = lines[1][6:]
	p.Timestamp = lines[2][11:]

	return nil
}

func (p Ping) Callback() {
	p.callback(p)
}

func (p *Ping) SetCallback(f func(ping Ping)) {
	p.callback = f
}

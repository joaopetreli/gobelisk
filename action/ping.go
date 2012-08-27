package action

import (
	"errors"
	"fmt"
)

type Ping struct {
	ActionName string
	PingResponse
}

func NewPing() *Ping {
	return &Ping{ActionName: "Ping"}
}

func (p *Ping) String() string {
	return fmt.Sprintf("ActionName: %s\n",
		p.ActionName)
}

func (p *Ping) Command() string {
	return fmt.Sprintf("Action: %s\n\n", p.ActionName)
}

type PingResponse struct {
	Success   bool
	Ping      string
	Timestamp string
}

func (pr *PingResponse) String() string {
	return fmt.Sprintf("Success: %b\nPing: %s\nTimestamp: %s\n",
		pr.Success, pr.Ping, pr.Timestamp)
}

func (pr *PingResponse) ResponseParser(r string) error {
	if len(r) != 59 {
		return errors.New("Invalid response.")
	}

	if r[10:17] == "Success" {
		pr.Success = true
	} else {
		pr.Success = false
	}

	pr.Ping = r[24:28]
	pr.Timestamp = r[40:57]

	return nil
}

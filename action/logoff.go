package action

import (
	"errors"
	"fmt"
)

type Logoff struct {
	ActionName string
	LogoffResponse
}

func NewLogoff() *Logoff {
	return &Logoff{ActionName: "Logoff"}
}

func (l *Logoff) String() string {
	return fmt.Sprintf("ActionName: %s\n", l.ActionName)
}

func (l *Logoff) Command() string {
	return fmt.Sprintf("Action: %s\n\n", l.ActionName)
}

type LogoffResponse struct {
	ExitMessage string
	Message     string
}

func (lr *LogoffResponse) String() string {
	return fmt.Sprintf("ExitMessage: %s\nMessage: %s\n",
		lr.ExitMessage, lr.Message)
}

func (lr *LogoffResponse) ResponseParser(r string) error {
	if len(r) != 53 {
		return errors.New("Invalid reponse.")
	}

	lr.ExitMessage = r[10:17]
	lr.Message = r[27:51]

	return nil
}

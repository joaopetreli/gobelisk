package manager

import (
	"bufio"
	"gobelisk/protocol/action"
	"net"
)

func sendQuery(conn net.Conn, act action.Action) error {
	if err := writeString(act.Query(), bufio.NewWriter(conn)); err != nil {
		return err
	}

	response, err := readBuffer(bufio.NewReader(conn))
	if err != nil {
		return err
	}

	act.Parse(response)
	act.Callback()

	return nil
}

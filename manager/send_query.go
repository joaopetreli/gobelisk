package manager

import (
	"bufio"
	"fmt"
	"gobelisk/protocol/action"
	"net"
)

func sendQuery(conn net.Conn, act action.Action) error {
	_, err := fmt.Fprint(conn, act.Query())
	if err != nil {
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

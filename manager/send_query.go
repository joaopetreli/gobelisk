package manager

import (
	"bufio"
	"fmt"
	"gobelisk/protocol/action"
	"net"
)

func SendQuery(conn net.Conn, act action.Action) error {
	_, err := fmt.Fprint(conn, act.Query())
	if err != nil {
		return err
	}

	response, err := readBuffer(bufio.NewReader(conn))
	if err != nil {
		return err
	}

	if err = act.Parse(response); err != nil {
		return err
	}

	act.Callback()

	return nil
}

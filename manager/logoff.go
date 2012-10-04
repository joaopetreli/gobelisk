package manager

import (
	"bufio"
	"fmt"
	"gobelisk/protocol/action"
	"net"
)

func Logoff(conn net.Conn) {
	logoff := action.NewLogoff()
	_, err := fmt.Fprint(conn, logoff.Query())
	if err != nil {
		fmt.Println(err)
	}

	response, err := readBuffer(bufio.NewReader(conn))
	if err != nil {
		fmt.Println(err)
		return
	}

	logoff.Parse(response)
	logoff.Callback()
}

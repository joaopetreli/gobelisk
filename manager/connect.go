package manager

import (
	"bufio"
	"fmt"
	"gobelisk/protocol/action"
	"gobelisk/protocol/event"
	"net"
)

// Warning: the caller should close the returned connection.
func Connect(host, port string, login *action.Login) (conn net.Conn, fullyBooted event.FullyBooted, err error) {
	conn, err = net.Dial("tcp", host+":"+port)
	if err != nil {
		return
	}

	// read the header
	header, err := read(conn)
	if err != nil {
		return
	}
	fmt.Print(header)

	// do login
	if err = writeString(login.Query(), bufio.NewWriter(conn)); err !=
	_, err = fmt.Fprint(conn, login.Query())
	if err != nil {
		return
	}

	// read login response
	response, err := readBuffer(bufio.NewReader(conn))
	if err != nil {
		return
	}

	// read the total response
	if len(response) == 55 {
		remainingResponse, err := readBuffer(bufio.NewReader(conn))
		if err != nil {
			return conn, event.FullyBooted{}, err
		}
		response += remainingResponse
	}

	// call loginresponse parser
	if err = login.Parse(response[:55]); err != nil {
		return
	}

	// call callback from login
	login.Callback()

	// read the dispatched event FullyBooted
	if fullyBooted.Parse(response[55:]); err != nil {
		return
	}

	return
}

func listenEvent(event event.Event) {

}

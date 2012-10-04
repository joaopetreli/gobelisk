package action

import (
	"fmt"
	"gobelisk/protocol"
	"strings"
)

type Login struct {
	Action   string
	Username string
	Secret   string
	callback func(Login)
	LoginResponse
}

type LoginResponse struct {
	Success     bool
	Response    string
	Message     string
	RawResponse string
}

func NewLogin(username, secret string) Login {
	var login Login

	login.Action = "Login"
	login.Username = username
	login.Secret = secret
	login.callback = func(l Login) {
		fmt.Print(l.RawResponse)
	}

	return login
}

func (l Login) Query() string {
	return fmt.Sprintf("Action: %s\nUsername: %s\nSecret: %s\r\n\r\n",
		l.Action, l.Username, l.Secret)
}

func (l *LoginResponse) Parse(response string) error {
	l.Success = false
	l.RawResponse = response

	responseLen := len(response)
	if responseLen != 55 && responseLen != 51 {
		return protocol.ErrInvalidResponse
	}

	if response[10:17] == "Success" {
		l.Success = true
	}

	lines := strings.Split(response, "\r\n")

	l.Response = lines[0][10:]
	l.Message = lines[1][9:]

	if l.Success == false {
		return protocol.ErrAuthenticationFailed
	}

	return nil
}

func (l Login) Callback() {
	l.callback(l)
}

func (l *Login) SetCallback(f func(login Login)) {
	l.callback = f
}

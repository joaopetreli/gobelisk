package action

import (
	"errors"
	"fmt"
)

type Login struct {
	ActionName string
	Username   string
	Secret     string
	callback   func()
	LoginResponse
}

func NewLogin(username, secret string) *Login {
	l := new(Login)
	l.ActionName = "Login"
	l.Username = username
	l.Secret = secret
	l.callback = func() {
		fmt.Println(l.LoginResponse)
	}

	return l
}

func (l *Login) String() string {
	return fmt.Sprintf("ActionName: %s\nUsername: %s\nSecret: %s\n",
		l.ActionName, l.Username, l.Secret)
}

func (l *Login) Command() string {
	return fmt.Sprintf("Action: Login\nUsername: %s\nSecret: %s\n\n",
		l.Username, l.Secret)
}

func (l *Login) Callback() func() {
	return l.callback
}

type LoginResponse struct {
	Success bool
	Message string
}

func (lr *LoginResponse) String() string {
	return fmt.Sprintf("Success: %b\nMessage: %s\n\n",
		lr.Success, lr.Message)
}

func (lr *LoginResponse) ResponseParser(r string) error {
	if len(r) != 52 {
		return errors.New("Invalid reponse.")
	}

	if r[10:17] == "Success" {
		lr.Success = true
	} else {
		lr.Success = false
	}

	lr.Message = r[27:50]

	return nil
}

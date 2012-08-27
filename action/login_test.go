package action

import (
	"fmt"
	"testing"
)

func TestNewLogin(t *testing.T) {
	l := NewLogin("user", "pass")
	if l.ActionName != "Login" {
		t.Fail()
	}

	if l.Username != "user" {
		t.Fail()
	}

	if l.Secret != "pass" {
		t.Fail()
	}
}

func TestLoginString(t *testing.T) {
	l := NewLogin("user", "pass")
	s := fmt.Sprintf("ActionName: Login\nUsername: user\nSecret: pass\n")
	if l.String() != s {
		t.Fail()
	}
}

func TestLoginCommand(t *testing.T) {
	l := NewLogin("user", "pass")
	s := fmt.Sprintf("Action: Login\nUsername: user\nSecret: pass\n\n")
	if l.Command() != s {
		t.Fail()
	}
}

func TestLoginResponseString(t *testing.T) {
	lr := new(LoginResponse)
	lr.Success = true
	lr.Message = "Authentication accepted"

	s := fmt.Sprintf("Success: %b\nMessage: %s\n\n",
		lr.Success, lr.Message)

	if lr.String() != s {
		t.Fail()
	}
}

func TestLoginResponseResponseWithInvalidResponse(t *testing.T) {
	lr := new(LoginResponse)
	err := lr.Response("Invalid response.")
	if err == nil {
		t.Fail()
	}
}
func TestLoginResponseResponseWithValidResponse(t *testing.T) {
	lr := new(LoginResponse)
	response := fmt.Sprintf("Response: Success\n" +
		"Message: Authentication accepted\n\n")
	err := lr.Response(response)
	if err != nil {
		t.Fail()
	}

	if lr.Success == false {
		t.Fail()
	}

	if lr.Message != "Authentication accepted" {
		t.Fail()
	}
}

package action

import (
	"fmt"
	"testing"
)

func TestNewLogoff(t *testing.T) {
	l := NewLogoff()
	if l.ActionName != "Logoff" {
		t.Fail()
	}
}

func TestLogoffString(t *testing.T) {
	l := NewLogoff()
	s := fmt.Sprintf("ActionName: %s\n", l.ActionName)

	if l.String() != s {
		t.Fail()
	}
}

func TestLogoffCommand(t *testing.T) {
	l := NewLogoff()
	s := fmt.Sprintf("Action: %s\n\n", l.ActionName)

	if l.Command() != s {
		t.Fail()
	}
}

func TestLogoffResponseString(t *testing.T) {
	lr := new(LogoffResponse)
	lr.ExitMessage = "Goodbye"
	lr.Message = "Thanks for all the fish."
	s := fmt.Sprintf("ExitMessage: %s\nMessage: %s\n",
		lr.ExitMessage, lr.Message)
	if lr.String() != s {
		t.Fail()
	}
}

func TestLogoffResponseInvalidResponse(t *testing.T) {
	lr := new(LogoffResponse)
	if err := lr.Response("Invalid response."); err == nil {
		t.Fail()
	}
}

func TestLogoffResponseValidResponse(t *testing.T) {
	lr := new(LogoffResponse)
	response := fmt.Sprintf("Response: Goodbye\n" +
		"Message: Thanks for all the fish.\n\n")
	if err := lr.Response(response); err != nil {
		t.Fail()
	}

	if lr.ExitMessage != "Goodbye" {
		t.Fail()
	}

	if lr.Message != "Thanks for all the fish." {
		t.Fail()
	}
}

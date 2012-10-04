package action

import (
	"gobelisk/protocol"
	"testing"
)

func TestLoginQuery(t *testing.T) {
	l := Login{
		Action:   "Login",
		Username: "username",
		Secret:   "secret",
	}

	query := "Action: Login\nUsername: username\nSecret: secret\r\n\r\n"
	if query != l.Query() {
		t.Errorf("Expected: '%s'\nbut got: '%s'", query, l.Query())
	}
}

func TestLoginResponseParseSuccess(t *testing.T) {
	response := "Response: Success\r\nMessage: Authentication accepted\r\n\r\n"
	l := NewLogin("username", "secret")

	if err := l.Parse(response); err != nil {
		t.Error(err)
	}

	if l.Success == false {
		t.Error("Should be a success login.")
	}

	if l.RawResponse != response {
		t.Error("login.RawResponse differs from manually generated response.")
	}

	if l.Response != "Success" {
		t.Errorf("Expected '%s'\nbut got: '%s'", "Success", l.Response)
	}

	if l.Message != "Authentication accepted" {
		t.Errorf("Expected '%s'\nbut got: '%s'", "Authentication accepted", l.Message)
	}
}

func TestLoginResponseParseFailure(t *testing.T) {
	response := "Response: Error\r\nMessage: Authentication failed\r\n\r\n"
	l := NewLogin("username", "secret")

	if err := l.Parse(response); err != protocol.ErrAuthenticationFailed {
		t.Error("Expected protocol.ErrAuthenticationFailed, but got", err)
	}

	if l.Success {
		t.Error("Should not be a success authentication.")
	}

	if l.RawResponse != response {
		t.Error("login.RawResponse differs from manually generated response.")
	}

	if l.Response != "Error" {
		t.Errorf("Expected '%s'\nbut got: '%s'", "Error", l.Response)
	}

	if l.Message != "Authentication failed" {
		t.Errorf("Expected '%s'\nbut got: '%s'", "Authentication failed", l.Message)
	}
}

func TestLoginResponseParseInvalidResponse(t *testing.T) {
	response := "Invalid response."
	l := NewLogin("username", "secret")

	if err := l.Parse(response); err != protocol.ErrInvalidResponse {
		t.Error("Expected protocol.ErrInvalidResponse, but got", err)
	}
}

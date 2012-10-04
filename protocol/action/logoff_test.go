package action

import (
	"gobelisk/protocol"
	"testing"
)

func TestLogoffQuery(t *testing.T) {
	l := Logoff{
		Action: "Logoff",
	}

	query := "Action: Logoff\r\n\r\n"
	if query != l.Query() {
		t.Error("Expected: '%s'\nbut gor: '%s'", query, l.Query())
	}
}

func TestLogoffResponseParseSuccess(t *testing.T) {
	l := NewLogoff()
	response := "Response: Goodbye\r\nMessage: Thanks for all the fish.\r\n\r\n"

	if err := l.Parse(response); err != nil {
		t.Error(err)
	}

	if response != l.RawResponse {
		t.Error("logoff.RawResponse differs from manually generated response.")
	}
}

func TestLogoffResponseParseInvalidResponse(t *testing.T) {
	l := NewLogoff()
	response := "Invalid response"

	if err := l.Parse(response); err != protocol.ErrInvalidResponse {
		t.Error("Expected protocol.ErrInvalidResponse, but got", err)
	}

	if response != l.RawResponse {
		t.Error("logoff.RawResponse differs from manually generated response.")
	}
}

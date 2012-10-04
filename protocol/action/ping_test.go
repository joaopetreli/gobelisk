package action

import (
	"gobelisk/protocol"
	"testing"
)

func TestPingQuery(t *testing.T) {
	p := NewPing()
	query := "Action: Ping\r\n\r\n"

	if p.Query() != query {
		t.Errorf("Expected: '%s'\nbut got: '%s'", query, p.Query())
	}
}

func TestPingResponseParseSuccess(t *testing.T) {
	response := "Response: Success\r\nPing: Pong\r\nTimestamp: 1349371187.324169\r\n\r\n"
	p := NewPing()

	if err := p.Parse(response); err != nil {
		t.Error(err)
	}

	if p.Response != "Success" {
		t.Error("Expected: 'Success', but got %s", p.Response)
	}

	if p.Ping != "Pong" {
		t.Error("Expected: 'Pong', but got %s", p.Ping)
	}

	if p.Timestamp != "1349371187.324169" {
		t.Error("Expected: '1349371187.324169', but got %s", p.Timestamp)
	}
}

func TestPingResponseParseInvalidResponse(t *testing.T) {
	response := "Invalid response."
	p := NewPing()

	if err := p.Parse(response); err != protocol.ErrInvalidResponse {
		t.Error("Expected protocol.ErrInvalidResponse, but got", err)
	}
}

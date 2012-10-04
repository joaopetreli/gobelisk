package action

import (
	// "gobelisk/protocol"
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

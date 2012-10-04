package action

import (
	"fmt"
	"testing"
)

func TestLoginQuery(t *testing.T) {
	l := Login{
		Action:   "Login",
		Username: "user",
		Secret:   "password",
	}

	query := fmt.Sprintf("Action: %s\nUsername: %s\nSecret: %s\r\n\r\n",
		"Login", "user", "password")
	if query != l.Query() {
		t.Errorf("Expected: '%s'\nbut got: '%s'", query, l.Query())
	}
}

package action

import (
	"fmt"
	"testing"
)

func TestNewPing(t *testing.T) {
	p := NewPing()
	if p.ActionName != "Ping" {
		t.Fail()
	}
}

func TestPingString(t *testing.T) {
	p := NewPing()
	s := fmt.Sprintf("ActionName: %s\n", p.ActionName)
	if p.String() != s {
		t.Fail()
	}
}

package protocol

import (
	"errors"
)

var (
	ErrInvalidResponse      = errors.New("Invalid response.")
	ErrAuthenticationFailed = errors.New("Authentication failed.")
)

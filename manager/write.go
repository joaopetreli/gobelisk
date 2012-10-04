package manager

import (
	"bufio"
	"fmt"
)

func writeString(s string, w *bufio.Writer) error {
	_, err := fmt.Fprint(w, s)
	if err != nil {
		return err
	}
	return nil
}

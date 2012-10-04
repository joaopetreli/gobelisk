package manager

import (
	"bufio"
	"io"
	"strings"
)

func read(rd io.Reader) (string, error) {
	var response []string
	for {
		rawPart := make([]byte, 27)
		_, err := rd.Read(rawPart)
		part := string(rawPart)
		response = append(response, part)

		if err != nil && err != io.EOF {
			return strings.Join(response, ""), err
		}

		if strings.HasSuffix(part, "\r\n") {
			break
		}
	}

	return strings.Join(response, ""), nil
}

func readBuffer(rd *bufio.Reader) (string, error) {
	var response []string
	for {
		rawPart := make([]byte, 500)
		bytesRead, err := rd.Read(rawPart)
		if bytesRead == 0 && err == io.EOF {
			break
		}

		part := string(rawPart[:bytesRead])
		response = append(response, part)
		if strings.HasSuffix(part, "\r\n\r\n") {
			break
		}

		if err != nil {
			return strings.Join(response, ""), err
		}
	}

	return strings.Join(response, ""), nil
}

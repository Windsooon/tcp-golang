/*
   Functions used by TCP server and TCP client
*/

package base

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

var portRangeError = errors.New("Port number should between 1 and 65535.")
var convertPortError = errors.New("Can't convert port to number")

// By default, the port number is 8080.
// If user provide a port number, this function
// will try to convert it to int.
func ReturnPort(input []string) (string, error) {
	// User 8080 by default
	if len(input) != 2 {
		port := "8080"
		return port, nil
	} else {
		err := ValidPort(input[1])
		if err != nil {
			return "", err
		} else {
			port := input[1]
			return port, nil
		}
	}
}

// This function check if we can convert the
// port to number between 1 and 65535
func ValidPort(port string) error {
	num, err := strconv.Atoi(port)
	if err == nil {
		if 1 <= num && num <= 65535 {
			return nil
		} else {
			return portRangeError
		}
	}
	return convertPortError
}

// Send Hello
func FormatMessage(message string) string {
	if len(message) == 0 {
		return fmt.Sprintf("<name> should not be empty." + "\n")
	}
	runes := []rune(message)
	for _, r := range runes {
		if !unicode.IsPrint(r) {
			return fmt.Sprintf("<name> is not printable." + "\n")
		}
	}
	return fmt.Sprintf("Hello " + message + "!" + "\n")
}

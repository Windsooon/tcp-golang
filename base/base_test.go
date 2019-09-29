package base

import (
	"testing"
)

func TestValidPort(t *testing.T) {
	var ports = []struct {
		portNumber string
		err        error
	}{
		{"1", nil},
		{"100", nil},
		{"65535", nil},
		{"70000", portRangeError},
		{"foo", convertPortError},
	}
	for _, v := range ports {
		t.Run(v.portNumber, func(t *testing.T) {
			err := ValidPort(v.portNumber)
			if err != v.err {
				t.Fatalf("expected: %v, got: %v", err, v.err)
			}
		})
	}
}

func TestFormatMessage(t *testing.T) {
	var messages = []struct {
		msg string
		res string
	}{
		{"foo", "Hello foo!\n"},
		{"bar", "Hello bar!\n"},
		{"100", "Hello 100!\n"},
		{"", "<name> should not be empty.\n"},
	}
	for _, v := range messages {
		t.Run(v.msg, func(t *testing.T) {
			val := FormatMessage(v.msg)
			if val != v.res {
				t.Fatalf("expected: %v, got: %v", val, v.res)
			}
		})
	}
}

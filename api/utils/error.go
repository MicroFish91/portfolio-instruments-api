package utils

import "regexp"

func redactError(err string) string {
	switch {
	case regexp.MustCompile(`(?i)failed[\s]to[\s]connect`).Match([]byte(err)) ||
		regexp.MustCompile(`(?i)connection[\s]refused`).Match([]byte(err)):
		return "internal error: could not connect to requisite database service"
	default:
		return err
	}
}

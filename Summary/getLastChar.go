package Summary

import (
	"errors"
	"strings"
)

//getLastChar return a string with the last given number of characters
func getLastChar(summary string, char int) (string, error) {
	if len(summary) >= char {
		return strings.ToLower(summary[len(summary)-char:]), nil
	} else {
		return "", errors.New("the summary is too small")
	}
}

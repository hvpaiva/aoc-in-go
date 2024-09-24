package parser

import (
	"strings"
)

// ToLines splits the input into lines and removes any leading or trailing whitespace from each line
// and empty last line if exists
func ToLines(input string) []string {
	input = Normalize(input)

	lines := strings.Split(input, "\n")

	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}

	return lines
}

// Normalize removes any leading or trailing whitespace from the input
func Normalize(input string) string {
	return strings.TrimSpace(input)
}

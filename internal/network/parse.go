package network

import "strings"

// ParseCliOutput takes CLI output and returns a slice of strings.  Each item
// in the slice represents a single line of output
func ParseCliOutput(output string) []string {
	retval := make([]string, 0)

	if len(strings.TrimSpace(output)) > 0 {
		retval = strings.Split(output, "\n")
	}

	return retval
}

// ParseCliOutputLine parses an output line and parses into a slice based on tokens in the string
func ParseCliOutputLine(src string) []string {
	fields := []string{}
	separator := ':'

	var currentField strings.Builder
	escaping := false

	for _, char := range src {
		if char == '\\' && !escaping {
			escaping = true
			continue
		}

		if char == separator && !escaping {
			fields = append(fields, currentField.String())
			currentField.Reset()
		} else {
			currentField.WriteRune(char)
			escaping = false
		}
	}

	// Append the last field
	fields = append(fields, currentField.String())

	return fields
}

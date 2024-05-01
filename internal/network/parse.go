package network

import "strings"

func parseCliOutput(src string) []string {
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

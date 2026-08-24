package utils

import "strconv"

func StringToBool(value string, defaultValue bool) bool {
	if value == "" {
		return defaultValue
	}

	parsed, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}

	return parsed
}

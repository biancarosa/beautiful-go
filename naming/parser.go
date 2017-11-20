package naming

import "strings"

func parse(s string) string {
	return strings.Replace(s, " ", "", -1)
}

func removeSpaces(s string) string {
	return strings.Replace(s, " ", "", -1)
}

func remove(s string, char string) string {
	return strings.Replace(s, char, "", -1)
}

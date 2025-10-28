package main

import "strings"

func myJoin(sep string, str ...string) string {
	var joinedString strings.Builder
	for _, s := range str {
		joinedString.WriteString(s)
		joinedString.WriteString(sep)
	}
	return joinedString.String()
}

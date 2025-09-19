package exer5

import "strings"

func expand(s string, f func(string) string) string {
	stringSplit := strings.Split(s, " ")
	var madenString strings.Builder
	for _, word := range stringSplit {
		isHashtag := strings.HasPrefix(word, "$")
		if isHashtag {
			madenString.WriteString(f(word[1:]))
			madenString.WriteString(" ")
		} else {
			madenString.WriteString(word)
			madenString.WriteString(" ")
		}
	}
	return madenString.String()
}

func myReplacementFunction(word string) string {
	return word
}

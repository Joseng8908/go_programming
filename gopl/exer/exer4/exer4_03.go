package main

import (
	"unicode"
	"unicode/utf8"
)

//4.3 ~ 4.7

//4.3
func reverseArrPtr(ptr *[5]int) {
	for i, j := 0, len(ptr) - 1; i < j; i, j = i+1, j-1{
		(*ptr)[i], (*ptr)[j] = (*ptr)[j], (*ptr)[i]
	}
}

func reverseSlice(s []int) {
	for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 2{
		s[i], s[j] = s[j], s[i]
	}
}

//4.4
func rotate(s []int, i int) {
	reverseSlice(s[:i])
	reverseSlice(s[i:])
	reverseSlice(s)	
}

//4.5
func removeNearString(strings []string) {

	if len(strings) == 0 {
		return 
	}

	i := 0
	for j := 1; j < len(strings); j++ {
		if strings[j] != strings[i] {
			i++
			strings[i] = strings[j]
		}
	}
}

//4.6
func compoundSpace(s []byte) {
	if len(s) == 0{
		return
	}

	write := 0
	inSpace := false

	for read := 0; read < len(s); {

		r, size := utf8.DecodeRune(s[read:])

		isSpace := unicode.IsSpace(r)	

		if !isSpace {
			if inSpace {
				s[write] = ' '
			} 
			utf8.EncodeRune(s[write:], r)
			write += size
			inSpace = false
		} else {
			if !inSpace {
				s[write] = ' '
				write++
				inSpace = true
			}
		}

		read += size
	}

}

//4.7
func reverseEncodedUTF8(s []byte) {

	for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1{
		s[i], s[j] = s[j], s[i]
	}
}



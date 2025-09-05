package main

import (
	"unicode"
	"fmt"
	"bufio"
	"io"
	"os"
	"unicode/utf8"
)


func charcount() {
	counts := make(map[rune]int)
	unikind := make(map[string]int)
	var utflen[utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			unikind["L"]++
		} else if unicode.IsPunct(r){
			unikind["P"]++
		} else if unicode.IsNumber(r){
			unikind["N"]++
		} else {
			
		}

		counts[r]++
		utflen[n]++
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Printf("\nkind\tcount\n")
	for k, v := range unikind {
		fmt.Printf("%s\t%d", k, v)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

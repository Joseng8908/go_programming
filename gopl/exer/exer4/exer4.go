package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {

	hashFlag256 := flag.Bool("sha256", false, "hash algorithm (ex.sha256)")
	hashFlag512 := flag.Bool("sha512", false, "hash algorithm (ex.sha512)")

	flag.Parse()
	args := flag.Args()

	if len(args) != 2 {
		fmt.Println("usage: go run exer4.go sha256|sha512 <string1> <string2>")
		return
	}

	if *hashFlag256 {
		c1 := sha256.Sum256([]byte(args[0]))
		c2 := sha256.Sum256([]byte(args[1]))
		fmt.Println(howManyDiff32(c1, c2))
	} else if *hashFlag512 {
		c1 := sha512.Sum512([]byte(args[0]))
		c2 := sha512.Sum512([]byte(args[1]))
		fmt.Println(howManyDiff64(c1, c2))
	} else {
		fmt.Println("please write flag")
	}
}

func howManyDiff32(c1, c2 [32]byte) int {
	sum := 0
	for i := 0; i < 32; i++ {
		result := c1[i] ^ c2[i]
		sum += PopCount(result)
	}
	return sum

}

func howManyDiff64(c1, c2 [64]byte) int {
	sum := 0
	for i := 0; i < 64; i++ {
		result := c1[i] ^ c2[i]
		sum += PopCount(result)
	}
	return sum

}

func PopCount(b byte) int {
	count := 0
	for b != 0 {
		b &= b - 1
		count++
	}
	return count
}

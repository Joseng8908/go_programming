package main

import (
	"bufio"
	"log"
	"os"
	"fmt"
)

func wordfreq() {
	countword := make(map[string]int)
	file, err := os.Open("hi.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		countword[word]++
	}

	fmt.Printf("word\tcount")
	for k, v := range countword{
		fmt.Printf("\n%s\t%d",k, v)
	}

	defer file.Close()
}

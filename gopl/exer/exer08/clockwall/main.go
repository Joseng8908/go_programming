package main

import (
	"log"
	"net"
	"os"
	"strings"
	"fmt" 
	"bufio"
)

func main() {
	for _, arg := range os.Args[1:] {
		parts := strings.Split(arg, "=")
		name := parts[0]
		address := parts[1]
		
		conn, err := net.Dial("tcp", address)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go func(n string, c net.Conn) {
			defer c.Close()
			watch(n, c)
		} (name, conn)
	}
	for{}
}
func watch(name string, r net.Conn) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Printf("%s: %s\n", name, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Printf("%s 에러: %v", name, err)
	}
}


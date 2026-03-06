package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	portPtr := flag.Int("port", 8000, "integer")
	flag.Parse()
	portNum := *portPtr

	address := fmt.Sprintf("localhost:%d", portNum)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close() 
	io.WriteString(c, "Welcome to Server\n")

	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		input := scanner.Text()
		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}
		command := args[0]
		switch command {
		case "cd":
			changeDir(args[1], c)
		case "ls":
			list(c)
		case "get":
			get(args[1], c)
		case "close":
			return
		default:
			io.WriteString(c, "unknown command")
		}
	}
}

func changeDir(dirName string, conn net.Conn) {
	os.Chdir(dirName)
}

func list(conn net.Conn) {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		io.WriteString(conn, file.Name() + "\n")
	}
}

func get(fileName string, conn net.Conn) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)	
	}
	defer file.Close()
	io.Copy(conn, file)
}


package main

import (
	"log"
	"net"
	"os"
)

const sockAddr = "/tmp/echo.sock"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("must pass more than two arugments")
	}

	conn, err := net.Dial("unix", sockAddr)
	if err != nil {
		log.Fatal("failed to dial: ", err)
	}

	if _, err := conn.Write([]byte(os.Args[1])); err != nil {
		log.Fatal("failed to write to socket", err)
	}

	b := make([]byte, 1024)
	if _, err := conn.Read(b); err != nil {
		log.Fatal("failed to read from socket", err)
	}
	log.Println("response: ", string(b))
}

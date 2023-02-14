package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const sockAddr = "/tmp/echo.sock"

func main() {
	socket, err := net.Listen("unix", sockAddr)
	if err != nil {
		log.Fatal(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Remove(sockAddr)
		os.Exit(1)
	}()

	for {
		conn, err := socket.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(conn net.Conn) {
			defer conn.Close()

			buf := make([]byte, 4096)

			n, err := conn.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("msg: ", string(buf))

			_, err = conn.Write(buf[:n])
			if err != nil {
				log.Fatal(err)
			}
		}(conn)
	}
}

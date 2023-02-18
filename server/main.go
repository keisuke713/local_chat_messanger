package main

import (
	"log"
	"os"
	"syscall"
)

const (
	serverAddr = "/tmp/echo.sock"
)

func main() {
	var (
		sockfd int
		err    error
	)

	err = os.Remove(serverAddr)
	if err != nil {
		log.Fatal("failed to close file due to ", err)
	}

	sockfd, err = syscall.Socket(syscall.AF_UNIX, syscall.SOCK_STREAM, 0)
	if err != nil {
		log.Fatal("failed to socket due to ", err)
	}

	sockaddr := &syscall.SockaddrUnix{
		Name: serverAddr,
	}
	err = syscall.Bind(sockfd, sockaddr)
	if err != nil {
		log.Fatal("failed to bind due to ", err)
	}

	err = syscall.Listen(sockfd, 1)
	if err != nil {
		log.Fatal("failed to listen due to ", err)
	}
	log.Println("start")

	for {
		// sa will not be used
		nfd, _, err := syscall.Accept(sockfd)
		if err != nil {
			log.Fatal("failed to accept due to ", err)
		}

		buf := make([]byte, 1024)
		n, err := syscall.Read(nfd, buf)
		if err != nil {
			log.Fatal("failed to read due to ", err)
		}
		log.Println("Received message which is ", string(buf[:n]))

		_, err = syscall.Write(nfd, []byte("success"))
		if err != nil {
			log.Fatal("faield to send due to ", err)
		}

		err = syscall.Close(nfd)
		if err != nil {
			log.Fatal("failed to close due to ", err)
		}
		log.Println("close")
	}
}

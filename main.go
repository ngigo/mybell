package main

//this is bell messenger API

import (
	"fmt"
	"log"
	"net"
	"os"
)

const (
	connHost = "localhost"
	connPort = "8000"
	connType = "tcp"
)

func main() {
	fmt.Println("Starting " + connType + " server on " + connHost + ":" + connPort + ".....")
	l, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("Client Connected!")

		fmt.Println("Client " + c.RemoteAddr().String() + " connected.")

	}

}

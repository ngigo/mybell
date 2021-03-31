package main

//this is bell messenger API

import (
	"bufio"
	"encoding/json"
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

func HandleJson(c net.Conn) {
	d := json.NewDecoder(c)

	var msg string

	err := d.Decode(&msg)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
	c.Close()

}

func handleConnection(conn net.Conn) {
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')

	if err != nil {
		fmt.Println("Client left.")
		conn.Close()
		return
	}

	log.Println("Client message:", string(buffer[:len(buffer)-1]))

	conn.Write(buffer)

	handleConnection(conn)
}

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

		go handleConnection(c)
	}

}

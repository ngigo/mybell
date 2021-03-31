package main

//this is bell messenger API

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

const (
	connHost = "localhost"
	connPort = "8000"
	connType = "tcp"
)

type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

//HandleJson reads(decode) json data
func HandleJson(c net.Conn, wg *sync.WaitGroup) {
	d := json.NewDecoder(c)

	var msg Message

	err := d.Decode(&msg)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg)
	wg.Done()
	c.Close()

}

func main() {
	var wg sync.WaitGroup

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

		wg.Add(1)
		go HandleJson(c, &wg)
		wg.Wait()

	}
}

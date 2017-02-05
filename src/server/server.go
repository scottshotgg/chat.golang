package main

import (
	"fmt"
	"net"
	"time"
	"github.com/satori/go.uuid"
)

var (
	listener 	net.Listener
	connChan	chan client	
)

type client struct {
	conn 		net.Conn
	connTime	time.Time
	uuid 		uuid.UUID
}

func accept() {
	for {
		conn, _ := listener.Accept()

		client := client{conn: conn, connTime: time.Now(), uuid: uuid.NewV4()}
		connChan <- client
	}
}

func printClient() {
	for {
		client:= <-connChan
		fmt.Println("Connection accepted from", client.conn.RemoteAddr(), "at", client.connTime, "with uuid:", client.uuid)
	}
}

func main() {
	fmt.Println("I am the server")
	listener, _ = net.Listen("tcp", ":8080")

	connChan = make(chan client) 

	go accept()

	go printClient()

	for {

	}

}
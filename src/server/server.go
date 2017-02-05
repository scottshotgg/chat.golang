package main

import (
	"fmt"
	"net"
	"time"
	"bufio"
	"github.com/satori/go.uuid"
)

var (
	listener 	net.Listener
	connChan	chan client	
	//readyChan	chan client	
	clientMap	map[uuid.UUID] net.Addr
)

type client struct {
	conn 		net.Conn
	connTime	time.Time
	uuid 		uuid.UUID
	reader 		*bufio.Scanner
	writer		*bufio.Writer
}

func (c client) Read() string {
	return c.reader.Text()
}

// func (c client) Write() {
// 	return 
// }

//func readClient() {

//}

func accept() {
	for {
		conn, _ := listener.Accept()

		client := client{	conn: conn, 
							connTime: time.Now(), 
							uuid: uuid.NewV4(),
							reader: bufio.NewScanner(bufio.NewReader(conn)),
							 writer: bufio.NewWriter(conn)}

		connChan <- client
	}
}

// func listenToClient() {
// 	for {

// 	}
// }

func processClient() {
	for {
		client:= <-connChan
		fmt.Println("Connection accepted from", client.conn.RemoteAddr(), "at", client.connTime, "with uuid:", client.uuid)

		// Insert the client
		clientMap[client.uuid] = client.conn.RemoteAddr()

		// Start client listener
		//go listenToClient()
	}
}

func main() {
	fmt.Println("I am the server")
	listener, _ = net.Listen("tcp", ":8080")

	connChan = make(chan client) 

	go accept()
	go processClient()

	for {

	}

}
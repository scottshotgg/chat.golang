package main

import (
	"fmt"
	"net"
	"time"
	"bufio"
	"strings"
	"github.com/satori/go.uuid"
)

type client struct {
	conn 		net.Conn
	connTime	time.Time
	uuid 		uuid.UUID
	reader 		*bufio.Reader
	writer		*bufio.Writer
}

var (
	listener 		net.Listener
	connChan		chan client
	keepListeningChan	chan int	
	//readyChan	chan client	
	// not sure if this should be uuid
	activeClientMap		map[uuid.UUID] client
	inActiveClientMap	map[uuid.UUID] client
)

// will do this later
// this uses the C preprocessor in the run script
func print(line string) {
	//#ifdef debug
	fmt.Print(line)
	//#endif
}

func (c client) Read() (string, int) {
	fmt.Println("waiting for something...")
	line, err := c.reader.ReadString('\n')
	fmt.Println("got something!", line, err)
	errInt := 1

	if err != nil { 
		fmt.Println(err.Error())
		// Put these here for now
		c.conn.Close()
		errInt = 0
	}

	fmt.Println("got here...")
	//keepListeningChan <- errInt
	return strings.TrimSuffix(line, "\n"), errInt
}

// can make a channel for each one waiting for stuff to send
func (c client) Write(line string) {
	c.writer.WriteString(line + "\n")
	c.writer.Flush()
}

func (c client) WriteAll(line string) {
	for _, value := range activeClientMap {	
		value.writer.WriteString(fmt.Sprintln(c.conn.RemoteAddr(), " :: ", time.Now().Format("Mon Jan _2 15:04:05 2006"), "::- ", line))
		value.writer.Flush()
	}
}

func (c client) PrintOut(line string) {
	fmt.Println(c.conn.RemoteAddr(), " :: ", time.Now().Format("Mon Jan _2 15:04:05 2006"), "::- ", line)
}

func accept() {
	for {
		// Need to do error handling here
		conn, _ := listener.Accept()

		client := client{		conn: conn, 
					connTime: time.Now(), 
					uuid: uuid.NewV4(),
					//reader: bufio.NewScanner(bufio.NewReader(conn)),
					 // OS 
					 // username
					 // local hash cookie
					 reader: bufio.NewReader(conn),
					 writer: bufio.NewWriter(conn)}

		// Maybe we should make a function called startChans or something
		connChan <- client
		//keepListeningChan <- 1
	}
}

func listenToClient(client client) {
	errInt := 1
	var line string

	for {
		switch  errInt {
			case 1:
				line, errInt = client.Read()
				fmt.Println(line)
				client.PrintOut(line)
				fmt.Println(strings.ToUpper(line))

				// this could be for the check with the client before we send it out
				//client.Write(strings.ToUpper(line))
				
				// just consider everything for now a server wide message
				// later make rooms, chats, and a server "objects" holding everything
				client.WriteAll(strings.ToUpper(line))
			case 0: 
				client.PrintOut("Client closed the connection")
				return
			default:
				fmt.Println("percolation is part of the water cycle i guess")
		}
	}
}

func processClient() {
	for {
		client:= <-connChan
		fmt.Println("Connection accepted from", client.conn.RemoteAddr(), "at", client.connTime, "with uuid:", client.uuid)

		// Insert the client
		//activeClientMap[client.uuid] = client.conn.RemoteAddr()
		activeClientMap[client.uuid] = client

		// Start client listener
		go listenToClient(client)
	}
}

func makeVars() {
 	activeClientMap = make(map[uuid.UUID] client)
	connChan = make(chan client) 
	keepListeningChan = make(chan int) 
}

func main() {
	// Try to make an auto builder for this using reflect
	print(fmt.Sprintf("%s %d\n", "debug sprintf test", 1))

	fmt.Println("I am the server")
	listener, _ = net.Listen("tcp", ":8080")

 	activeClientMap = make(map[uuid.UUID] client)
	connChan = make(chan client) 
	keepListeningChan = make(chan int) 

	go accept()
	go processClient()

	for {

	}

}
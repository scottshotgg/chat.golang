package main

import (
	"fmt"
	"net"
	"time"
	"bufio"
	"strings"
	"database/sql"

	"github.com/satori/go.uuid"
	_ "github.com/mattn/go-sqlite3"
)

var (
	listener 		net.Listener

	connChan		chan Client
	keepListeningChan	chan int
	commandChan		chan string

	activeClientMap	map[uuid.UUID] Client
	inActiveClientMap	map[uuid.UUID] Client
	DB                    		*sql.DB
	//readyChan	chan client	
	// not sure if this should be uuid
)

type Client struct {
	conn 		net.Conn
	connTime	time.Time
	uuid 		uuid.UUID
	reader 		*bufio.Reader
	writer		*bufio.Writer
	//id
	//os
}

func (c Client) Read() (string, int) {
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
func (c Client) Write(line string) {
	c.writer.WriteString(line + "\n")
	c.writer.Flush()
}

// func (c Client) WriteAll(line string, activeClientMap map[uuid.UUID] Client) {
func (c Client) WriteAll(line string) {
	for _, value := range activeClientMap {	
		value.writer.WriteString(fmt.Sprintln(c.conn.RemoteAddr(), "::", time.Now().Format("Mon Jan _2 15:04:05 2006"), "::-", line))
		value.writer.Flush()
	}
}

func (c Client) PrintOut(line string) {
	fmt.Println(c.conn.RemoteAddr(), "::", time.Now().Format("Mon Jan _2 15:04:05 2006"), "::-", line)
}

// will do this later
// this uses the C preprocessor in the run script
func print(line string) {
	//#ifdef debug
	fmt.Print(line)
	//#endif
}

func makeVars() {
 	activeClientMap = 	make(map[uuid.UUID] Client)
	connChan = 		make(chan Client) 
	keepListeningChan = 	make(chan int) 
	commandChan = 	make(chan string, 10)
}

func parseCommand(client Client) {

	for {
		command := strings.Fields(<-commandChan)
		fmt.Println(command)

		switch command[0] {
			case "save":
				fmt.Println("save")
				// this will go to a db
			case "workout":
				// workout [day/date; relative;absolute] excersise sets reps weight
				fmt.Println("workout")
				// later this will use a username and passwd
				// use this later: https://github.com/tealeg/xlsx

				// submit this to a database thread
			default:
				fmt.Println("chat")
				client.WriteAll(strings.Join(command[1:], " "))
		}
	}
}

func listenToClient(client Client) {
	errInt := 1
	var line string

	go parseCommand(client)

	for {
		switch  errInt {
			case 1:
				line, errInt = client.Read()
				fmt.Println(line)
				client.PrintOut(line)
				fmt.Println(strings.ToUpper(line))

				commandChan <- line



			case 0: 
				client.PrintOut("Client closed the connection")
				return
			default:
				fmt.Println("Percolation is part of the water cycle i guess")
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


func accept() {
	for {
		// Need to do error handling here
		conn, _ := listener.Accept()

		client := Client{	conn: conn, 
					connTime: time.Now(), 
					uuid: uuid.NewV4(),
					//reader: bufio.NewScanner(bufio.NewReader(conn)),
					 // OS 
					 // username
					 // local hash cookie
					 reader: bufio.NewReader(conn),
					 writer: bufio.NewWriter(conn)}

		connChan <- client
	}
}

// Start here; climb up for readability
func main() {
	DB, err := sql.Open("sqlite3", "workout.db")
	fmt.Println(DB, err)

	// Try to make an auto builder for this using reflect
	print(fmt.Sprintf("%s %d\n", "debug sprintf test", 1))

	fmt.Println("I am the server")
	listener, _ = net.Listen("tcp", ":8080")

	makeVars()

	go accept()
	go processClient()

	// Make this go away somehow ??
	for {}
}
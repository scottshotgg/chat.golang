package client

import (
	"fmt"
	"net"
	"time"
	"bufio"
	"strings"
	"github.com/satori/go.uuid"
)

type Client struct {
	conn 		net.Conn
	connTime	time.Time
	uuid 		uuid.UUID
	reader 		*bufio.Reader
	writer		*bufio.Writer
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

func (c Client) WriteAll(line string, activeClientMap map[uuid.UUID] Client) {
	for _, value := range activeClientMap {	
		value.writer.WriteString(fmt.Sprintln(c.conn.RemoteAddr(), " :: ", time.Now().Format("Mon Jan _2 15:04:05 2006"), "::- ", line))
		value.writer.Flush()
	}
}
package main
import (
 "fmt"
 "net"
 "time"
 "bufio"
 "strings"
 "log"
 "github.com/satori/go.uuid"
)
var (
 listener net.Listener
 connChan chan client
 keepListeningChan chan int
 activeClientMap map[uuid.UUID] net.Addr
 inActiveClientMap map[uuid.UUID] net.Addr
)
type client struct {
 conn net.Conn
 connTime time.Time
 uuid uuid.UUID
 reader *bufio.Reader
 writer *bufio.Writer
}
func print(line string) {
}
func (c client) Read() (string, int) {
 print("waiting for something...")
 line, err := c.reader.ReadString('\n')
 print("got something!", line, err)
 errInt := 1
 if err != nil {
  fmt.Println(err.Error())
  c.conn.Close()
  errInt = 0
 }
 print("got here...")
 return strings.TrimSuffix(line, "\n"), errInt
}
func (c client) Write(line string) {
 c.writer.WriteString(line + "\n")
 c.writer.Flush()
}
func (c client) PrintOut(line string) {
 log.Println(c.conn.RemoteAddr(), " :: ", time.Now().Format("Mon Jan _2 15:04:05 2006"), line)
}
func accept() {
 for {
  conn, _ := listener.Accept()
  client := client{ conn: conn,
     connTime: time.Now(),
     uuid: uuid.NewV4(),
      reader: bufio.NewReader(conn),
      writer: bufio.NewWriter(conn)}
  connChan <- client
 }
}
func listenToClient(client client) {
 errInt := 1
 var line string
 for {
  switch errInt {
   case 1:
    line, errInt = client.Read()
    print(line)
    client.PrintOut(line)
    print(strings.ToUpper(line))
    client.Write(strings.ToUpper(line))
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
  log.Println("Connection accepted from", client.conn.RemoteAddr(), "at", client.connTime, "with uuid:", client.uuid)
  activeClientMap[client.uuid] = client.conn.RemoteAddr()
  go listenToClient(client)
 }
}
func makeVars() {
  activeClientMap = make(map[uuid.UUID] net.Addr)
 connChan = make(chan client)
 keepListeningChan = make(chan int)
}
func main() {
 print(fmt.Sprintf("%s %d\n", "debug sprintf test", 1))
 print("I am the server")
 listener, _ = net.Listen("tcp", ":8080")
  activeClientMap = make(map[uuid.UUID] net.Addr)
 connChan = make(chan client)
 keepListeningChan = make(chan int)
 go accept()
 go processClient()
 for {
 }
}

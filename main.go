package main

import "net"
import "fmt"
//import "bufio"
import "os"

const (
	CONN_HOST = "localhost"
	CONN_PORT = "4040"
	CONN_TYPE = "tcp"
	VERSION = "Prerun v0.1a"
)

func main() {
	// connect to this socket
	conn, err := net.Dial("tcp", "127.0.0.1:4040")
	if err != nil {
		fmt.Errorf("Connection failed: %q", err.Error())
		os.Exit(1)
	}
	buf_receive := make([]byte, 1024)
	//buf_send := make([]byte, 1024)
	for {
		// read in input from stdin
		//reader := bufio.NewReader(os.Stdin)
		///fmt.Print("Text to send: ")
		//text, _ := reader.ReadString('\n')

		// send to socket
		conn.Write([]byte(VERSION))
		// listen for reply
		reqLen, err := conn.Read(buf_receive)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			break
		} else if reqLen != 0 {
			fmt.Printf("Received: %q", buf_receive[:reqLen])
			break
		}

	}
}
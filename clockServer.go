// Clock2 is a concurrent TCP server that periodically writes the time.
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n")) // write the time
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second) // wait 1s
	}
}

func main() {
	if len(os.Args) < 3 { // check if the number of arguments is correct
		fmt.Println("Invlid number of arguments.")
		return
	}

	if os.Args[1] != "-port" { // check if the first argument is -port
		fmt.Println("Invalid flag.")
		return
	}

	// Listen on TCP port and run on argument 2
	l, err := net.Listen("tcp", "localhost:"+os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
}

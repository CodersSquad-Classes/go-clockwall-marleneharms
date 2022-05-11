package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func printTime(dst io.Writer, src io.Reader, city string) { // print the time dsr is the destination, src is the source, city is the city name
	s := bufio.NewScanner(src) // create a scanner
	for s.Scan() {             // scan the time
		fmt.Printf("%s\t: %s\n", city, s.Text()) // print the time
	}
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		arr := strings.Split(os.Args[i], "=") // split the argument

		if len(arr) != 2 { // check if the number of arguments is correct
			fmt.Println("Invalid format for arguments. Format: <cityName>=<portNumber>")
			return
		}

		conn, err := net.Dial("tcp", arr[1]) // connect to the port
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close() // close the connection when the function ends

		go printTime(os.Stdout, conn, arr[0]) // print the time concurrently
	}

	ch := make(chan int) // create a channel
	<-ch                 // wait for the channel to be closed
}

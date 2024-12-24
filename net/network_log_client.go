package main

import (
	"log"
	"net"
	"time"
)

// Need to set up the log server first
// For example, run the NetCat server in the background by the following command:
// $ nc -lk 1902
// -l: listen mode
// -k: Forces nc to stay listening for another connection after its current connection is completed.
// -u: Use UDP instead of the default option of TCP.
func main() {
	// conn, err := net.Dial("tcp", "localhost:1902")

	// use UDP to connect to the server for addressing the back-pressure problem
	timeout := 5 * time.Second
	conn, err := net.DialTimeout("udp", "localhost:1902", timeout)
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()

	log.Println("Connected to the server")

	f := log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC | log.Lshortfile
	logger := log.New(conn, "example: ", f)
	logger.Println("This is a regular message.")
	logger.Panicln("This is a panic.")
}

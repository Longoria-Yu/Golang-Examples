package main

import (
	"log"
	"net"
)

// Need to set up the log server first
// For example, run the NetCat server in the background by the following command:
// $ nc -lk 1902
// -l: listen mode
// -k: Forces nc to stay listening for another connection after its current connection is completed.
func main() {
	conn, err := net.Dial("tcp", "localhost:1902")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	log.Println("Connected to the server")

	f := log.Ldate | log.Ltime | log.Lshortfile
	logger := log.New(conn, "example: ", f)
	logger.Println("This is a regular message.")
	logger.Panicln("This is a panic.")
}

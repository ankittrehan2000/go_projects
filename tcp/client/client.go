package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	const name = "writetcp"
	log.SetPrefix(name + "\t") // set prefix for the standard logger

	// register command line flags
	port := flag.Int("p", 8080, "port to connect to")
	flag.Parse()

	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{Port: *port})
	if err != nil {
		log.Fatalf("error connecting to localhost:%d: %v", *port, err)
	}

	// close connection after function exits - a.k.a clean up
	defer conn.Close()

	// get a goroutine to read incoming lines from the server
	// TCP is duplex which means we can read and write at the same time
	// But that also means that we need to spawn a goroutine for the reading
	go func() {
		for connScanner := bufio.NewScanner(conn); connScanner.Scan(); {
			fmt.Printf("%s\n", connScanner.Text())

			if err := connScanner.Err(); err != nil {
				log.Fatalf("error reading from %s: %v", conn.RemoteAddr(), err)
			}
		}
	}()

	// read incoming lines from the stdin (CLI) and forward them to server
	for stdinScanner := bufio.NewScanner(os.Stdin); stdinScanner.Scan(); {
		log.Printf("sent: %s\n", stdinScanner.Text())
		// .Bytes returns bytes until the new line character
		if _, err := conn.Write(stdinScanner.Bytes()); err != nil {
			log.Fatalf("error writing to %s: %v", conn.RemoteAddr(), err)
		}
		if _, err := conn.Write([]byte("\n")); err != nil {
			log.Fatalf("error writing to %s: %v", conn.RemoteAddr(), err)
		}
		if stdinScanner.Err() != nil {
			log.Fatalf("error writing to %s: %v", conn.RemoteAddr(), err)
		}
	}
}

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

// echoUpper reads lines from r, uppercases them, and writes them to w.
func echoUpper(w io.Writer, r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		// note that scanner.Text() strips the newline character from the end of the line,
		// so we need to add it back in when we write to w.
		fmt.Fprintf(w, "%s\n", strings.ToUpper(line))
	}
	if err := scanner.Err(); err != nil {
		log.Printf("error: %s", err)
	}
}

func main() {
	const name = "tcpupperecho"
	log.SetPrefix(name + "\t")

	port := flag.Int("p", 8080, "port to listen on")
	flag.Parse()

	// Listen TCP creates a tcp connection by listening on a given address
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{Port: *port})

	if err != nil {
		panic(err)
	}

	defer listener.Close()
	log.Printf("listening at localhost: %s", listener.Addr())

	for {
		// loop forever to accept connections
		conn, err := listener.Accept()

		if err != nil {
			panic(err)
		}

		go echoUpper(conn, conn)
	}
}

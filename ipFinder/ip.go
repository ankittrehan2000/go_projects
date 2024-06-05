package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	log.Printf("%v", os.Args)
	if len(os.Args) != 2 {
		log.Fatalf("expected exactly one argument; got %d", len(os.Args)-1)
	}

	host := os.Args[1]
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatalf("lookup ip: %s %v", host, err)
	}

	for _, ip := range ips {
		if ip.To4() != nil {
			fmt.Println(ip)
			goto IPV6
		}
	}

	fmt.Printf("none \n")

IPV6:
	for _, ip := range ips {
		if ip.To4() == nil {
			fmt.Println(ip)
			return
		}
	}

	fmt.Printf("none \n")
}

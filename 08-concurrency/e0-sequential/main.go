package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	for i := 5200; i <= 6500; i++ {
		conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", i))
		if err != nil {
			log.Printf("%d CLOSED (%s)\n", i, err)
			continue
		}
		conn.Close()
		log.Printf("%d OPEN\n", i)
	}
}

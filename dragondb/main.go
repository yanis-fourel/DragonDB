package main

import (
	"dragon/store"
	"fmt"
	"io"
	"log"
	"net"
)

const port = 6969

func main() {
	store, err := store.New()
	if err != nil {
		log.Fatalln("Error creating store: ", err)
	}
	defer store.Close()

	l, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalln("Error listening to TCP port: ", err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalln("Error accepting connection: ", err)
		}

		go handleClient(conn, store)
	}
}

func handleClient(conn net.Conn, store *store.Store) {
	defer conn.Close()

	for {
		buf := make([]byte, 129)

		_, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Println("Error reading from TCP channel: ", err)
			return
		}

		cmd := string(buf)
		log.Println("Received command: ", cmd)
	}
}

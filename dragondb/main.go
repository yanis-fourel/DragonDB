package main

import (
	"dragon/store"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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

		cmdLen, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Println("Error reading from TCP channel: ", err)
			return
		}

		cmd := string(buf[:cmdLen])
		cmd = strings.TrimSuffix(cmd, "\n")

		key, value, isSetCmd := strings.Cut(cmd, "=")
		if isSetCmd {
			fmt.Printf("Key: '%s', Value: '%s'\n", key, value)
			store.Set(key, value)
			conn.Write([]byte("OK\n"))
		} else {
			fmt.Println("Getting key: ", key)
			value := store.Get(key)
			conn.Write([]byte(value + "\n"))
		}
	}
}

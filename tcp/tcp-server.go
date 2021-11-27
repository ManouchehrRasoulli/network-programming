package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// the server part
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		panic(err)
	}
	defer listener.Close() // close listening process after termination
	fmt.Println(listener.Addr())

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("got error (%v) on accepting connection\n", err)
		}
		defer conn.Close()
		conn.Write([]byte("some things are happening to me !!"))
	}()

	// the client part
	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		panic(fmt.Errorf("got error (%v) on connection", err))
	}

	arr := make([]byte, 100)
	conn.Read(arr)
	fmt.Println(string(arr))
	conn.Close()
	listener.Close()
}

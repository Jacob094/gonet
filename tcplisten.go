package main

import (
	"fmt"
	"log"
	"net"
)

func ListenToTcp(p string) {

	// second arg of net.Listen will be p
	t, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		log.Fatal(err)
	}

	defer t.Close()

	// This code is not optimal, but just wanting to test if I can accept TCP connections + read messages
	for {
		conn, err := t.Accept()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Connection made: ", conn.LocalAddr())

		x := make([]byte, 255)

		read, err := conn.Read(x)
		if err != nil {
			log.Fatal(err)
		}

		x = append(x, x[:read]...)
		fmt.Println(string(x))
		fmt.Println(x)
		conn.Close()
	}
}

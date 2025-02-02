package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Welcome to Tedis")

	l, err := net.Listen("tcp", "0.0.0.0:6379") // by defauly redis server is listen on 6379
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	_, err = l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
}

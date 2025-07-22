package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connection:", err)
		os.Exit(1)
	}
	defer conn.Close()

	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading from server:", err)
		os.Exit(1)
	}
	if message != "OK\n" {
		fmt.Printf("Want 'OK', got %s\n", message)
		os.Exit(1)
	}

	fmt.Println("Received from server: ", message)
}

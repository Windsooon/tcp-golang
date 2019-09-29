/*
   TCP server written in Golang by Windson Yang
*/
package main

import (
	"./base"
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

// Before Launching the TCP server
func beforeLaunching() {
	fmt.Println("This is a TCP server...")
}

// Start Launching the TCP server
func startLaunching(port string) (net.Listener, error) {
	fmt.Println("Launching Server at", port)
	listener, err := net.Listen("tcp", ":"+port)
	return listener, err
}

// Handle the client request
func handleConnection(conn net.Conn) {
	for {
		// Get the remote address of the client
		clientAddr := conn.RemoteAddr().String()
		bufferBytes, err := bufio.NewReader(conn).ReadBytes('\n')
		if err != nil {
			log.Println("Client from " + clientAddr + " left..")
			conn.Close()
			return
		}
		// Convert bytes from buffer to string
		message := string(bufferBytes)
		message = message[:len(message)-1]
		response := fmt.Sprintf(message + " from " + clientAddr)
		log.Println(response)
		// Format the return message
		returnMessage := base.FormatMessage(message)
		conn.Write([]byte(returnMessage))
	}
}

func main() {
	beforeLaunching()
	arguments := os.Args
	port, returnErr := base.ReturnPort(arguments)
	if returnErr != nil {
		log.Println(returnErr)
		os.Exit(0)
	}
	listener, startErr := startLaunching(port)
	// Failed to start the TCP server
	if startErr != nil {
		log.Println(startErr)
		os.Exit(0)
	}
	for {
		conn, acceptErr := listener.Accept()
		if acceptErr != nil {
			log.Println(acceptErr)
			os.Exit(0)
		}
		go handleConnection(conn)
	}
	listener.Close()
}

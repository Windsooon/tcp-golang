/*
   TCP client written in Golang by Windson Yang
*/
package main

import (
	"./base"
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	port, returnErr := base.ReturnPort(arguments)
	if returnErr != nil {
		log.Println(returnErr)
		os.Exit(0)
	}
	// Connect to port
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		if strings.TrimSpace(string(text)) == "Exit" {
			fmt.Println("TCP client exiting...")
			return
		}
		fmt.Fprintf(conn, text+"\n")
		// Print message from server
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print(">> " + message)
	}
}

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"hw6/config"
)

func main() {

	conn, _ := net.Dial(config.Network, config.UrlAddress)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}

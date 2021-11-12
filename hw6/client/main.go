package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"hw6/config"
)

func main() {

	// Подключаемся к сокету
	conn, _ := net.Dial(config.Network, config.UrlAddress)
	for {
		// Чтение входных данных от stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// Отправляем в socket
		fmt.Fprintf(conn, text)
		// Прослушиваем ответ
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}

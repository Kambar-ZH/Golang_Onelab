package main

import (
	"bufio"
	"fmt"
	"hw6/config"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func (s *Server) handleConnection(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
        message := scanner.Text()
		fmt.Println("Message Received:", message)
		number, err := strconv.Atoi(message)
		if err != nil {
            conn.Write([]byte("bad request: expected integer\n"))
            continue
		}
        fmt.Println("Waiting for heavy computation...")
        time.Sleep(5 * time.Second)
		squareNumber := number * number
		conn.Write([]byte(fmt.Sprintf("%d\n", squareNumber)))
        fmt.Println("The work is done!")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error:", err)
	}
}

type Server struct {
	listener net.Listener
	quit     chan interface{}
	wg       sync.WaitGroup
}

func NewServer(addr string) *Server {
	s := &Server{
		quit: make(chan interface{}),
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	s.listener = ln
	s.wg.Add(1)
	go s.serve()
	return s
}

func (s *Server) serve() {
	defer s.wg.Done()
    fmt.Println("start serving on 8081 port")
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			select {
			case <-s.quit:
                fmt.Println("Graceful termination")
				return
			default:
				log.Println("accept error", err)
			}
		} else {
			s.wg.Add(1)
			go func() {
				s.handleConnection(conn)
				s.wg.Done()
			}()
		}
	}

}

func (s *Server) Stop() {
	close(s.quit)
	s.listener.Close()
	s.wg.Wait()
}

func main() {
	s := NewServer(config.UrlAddress)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    s.wg.Add(1)
	go func() {
		<-c
        s.Stop()
		s.wg.Done()
	}()
    s.wg.Wait()
}

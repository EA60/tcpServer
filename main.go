package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

type MyServer struct {
	La string
	Ln net.Listener
	w  chan struct{}
}

func NewServer(a string) *MyServer {
	return &MyServer{
		La: a,
		w:  make(chan struct{}),
	}
}

func (s *MyServer) Start() error {
	ln, err := net.Listen("tcp", s.La)
	if err != nil {
		fmt.Println("error Start")
		log.Fatalln(err)
	}
	s.Ln = ln
	defer s.Ln.Close()

	go s.acceptLoop()

	<-s.w

	return nil
}

func (s *MyServer) acceptLoop() {
	for {
		conn, err := s.Ln.Accept()
		if err != nil {
			fmt.Println("error accept loop")
			log.Fatalln(err)
			continue
		}
		go s.readLoop(conn)
	}
}

func (s *MyServer) readLoop(conn net.Conn) {
	// go s.timeOut()

	defer conn.Close()
	b := make([]byte, 512)
	for {
		a, err := conn.Read(b)
		if err != nil || err == io.EOF {
			fmt.Println("error read loop")
			log.Fatalln(err)
		}
		message := b[:a]
		fmt.Println(string(message))
	}
}
func (s *MyServer) timeOut() {
	time.Sleep(time.Second * 10)
	s.w <- struct{}{}
}

func main() {
	server := NewServer(":8000")
	log.Fatal(server.Start())

}

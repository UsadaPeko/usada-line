package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	con, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   net.IPv4zero,
		Port: 1533,
	})
	if err != nil {
		log.Fatal(err)
	}
	lc := NewSimpleLineClient(con)
	lc.Start()

	select {}
}

type SimpleLineClient struct {
	con *net.TCPConn
}

func NewSimpleLineClient(con *net.TCPConn) *SimpleLineClient {
	return &SimpleLineClient{
		con: con,
	}
}

func (lc *SimpleLineClient) Hello() {
	hello := "hello"
	lc.SendMessage([]byte(hello))
	log.Println("HELLO~~!!")
}

func (lc *SimpleLineClient) SendMessage(message []byte) {
	n := len(message)
	rwn := 0
	for rwn != n {
		wn, err := lc.con.Write(message)
		if err != nil {
			log.Println(err)
			return
		}
		rwn += wn
	}
	log.Println("SENT MESSAGE!")
}

func (lc *SimpleLineClient) Input() []byte {
	message := make([]byte, 2048)
	_, err := fmt.Scanln(&message)
	if err != nil {
		log.Println(err)
	}
	return message
}

func (lc *SimpleLineClient) Start() {
	lc.Hello()

	go func() {
		for {
			lc.SendMessage(lc.Input())
		}
	}()
}

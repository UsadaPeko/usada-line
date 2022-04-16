package interfaces

import (
	"log"
	"net"
)

type TCPServer struct {
	listener *net.TCPListener
}

func NewTCPServer() *TCPServer {
	return &TCPServer{}
}

func (ts *TCPServer) Start() {
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{
		IP:   net.IPv4zero,
		Port: 1533,
	})
	if err != nil {
		log.Fatal(err)
	}
	ts.listener = listener

	go func() {
		for {
			ts.Connect()
		}
	}()
}

func (ts *TCPServer) Connect() {
	con, err := ts.listener.AcceptTCP()
	log.Println("TCP Connection is connected")

	err = con.SetKeepAlive(true)
	if err != nil {
		log.Println("connection set keep-alive failed")
	}
	cc := NewChatConnection(con)
	go cc.StartHandle()
}

func (ts *TCPServer) Shutdown() {

}

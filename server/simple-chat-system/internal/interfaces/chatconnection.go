package interfaces

import (
	"github.com/UsadaPeko/usadaline/simplechatsystem/internal/chatroom/domain"
	"github.com/UsadaPeko/usadaline/simplechatsystem/internal/dependencies"
	"io"
	"log"
	"net"
	"strings"
)

type ChatConnection struct {
	con  *net.TCPConn
	name string
}

func NewChatConnection(con *net.TCPConn) *ChatConnection {
	return &ChatConnection{con: con}
}

func (cc *ChatConnection) StartHandle() {
	cc.Hello()

	go func() {
		for {
			message, err := cc.Read()
			if err != nil {
				log.Printf("user(%v): connection end\n", cc.name)
				return
			}
			log.Printf("user(%v) -> %v\n", cc.name, string(message))
		}
	}()
}

func (cc *ChatConnection) Hello() {
	helloMessage, _ := cc.Read()
	hello := string(helloMessage)
	if strings.Compare(hello, "hello") == 0 {
		log.Printf("hello is failed. message is %v.\n", hello)
	}
	nameInput, _ := cc.Read()
	name := string(nameInput)
	dependencies.Use().UserInfoUseCases().CreateNewUser(name)
	dependencies.Use().ChatRoomUseCases().DetectNewUser(domain.NewUser(name))
	cc.name = name
}

func (cc *ChatConnection) Read() ([]byte, error) {
	read := make([]byte, 2048)
	for {
		localRead := make([]byte, 127)
		n, err := cc.con.Read(localRead)
		if err == io.EOF {
			return read, err
		}
		if err != nil {
			log.Println(err)
			return read, err
		}
		if n == 0 {
			log.Println("END")
			return read, nil
		}
		read = append(read, localRead...)

		if n < 127 {
			break
		}
	}
	return read, nil
}

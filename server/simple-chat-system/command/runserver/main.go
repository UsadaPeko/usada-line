package main

import "github.com/UsadaPeko/usadaline/simplechatsystem/internal/interfaces"

func main() {
	ts := interfaces.NewTCPServer()
	ts.Start()

	select {}
}

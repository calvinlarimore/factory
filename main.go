package main

import (
	"github.com/calvinlarimore/factory/server"
)

const (
	host = "localhost"
	port = 42069
)

func main() {
	server.StartServer(host, port)
}

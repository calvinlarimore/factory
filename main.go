package main

import (
	"github.com/calvinlarimore/factory/game"
)

const (
	host = "localhost"
	port = 42069
)

func main() {
	game.InitWorld()
	game.StartServer(host, port)
}

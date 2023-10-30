package main

import (
	"github.com/calvinlarimore/factory/game"
)

const (
	host = "0.0.0.0"
	port = 11111
)

func main() {
	game.InitWorld()
	game.StartServer(host, port)
}

package client

import (
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
)

func Middleware() wish.Middleware {
	return func(h ssh.Handler) ssh.Handler {
		return func(s ssh.Session) {
			wish.Println(s, "factory")
			h(s)
		}
	}
}

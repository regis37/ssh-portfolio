package server

import (
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	bm "github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"

	"github.com/charmbracelet/ssh"
	"github.com/regis37/ssh-portfolio/internal/ui"
)

const (
	host = "127.0.0.1"
	port = "23234"
)

func New() (*ssh.Server, error) {
	srv, err := wish.NewServer(
		wish.WithAddress(host+":"+port),
		wish.WithMiddleware(
			bm.Middleware(ui.TeaHandler),
			activeterm.Middleware(),
			logging.Middleware(),
		),
	)
	if err != nil {
		return nil, err
	}
	log.Info("SSH portfolio server created", "host", host, "port", port)
	return srv, nil
}

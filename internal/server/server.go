package server

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	bm "github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"

	"github.com/regis37/ssh-portfolio/internal/ui"
)

const defaultAddr = "127.0.0.1:23234"

func New() (*ssh.Server, error) {
	addr := os.Getenv("PORTFOLIO_ADDR")
	if addr == "" {
		addr = defaultAddr
	}

	opts := []ssh.Option{
		wish.WithAddress(addr),
		wish.WithMiddleware(
			bm.Middleware(ui.TeaHandler),
			activeterm.Middleware(),
			logging.Middleware(),
		),
	}

	hostKey := os.Getenv("PORTFOLIO_HOST_KEY")
	if hostKey != "" {
		opts = append(opts, wish.WithHostKeyPath(hostKey))
	}

	srv, err := wish.NewServer(opts...)
	if err != nil {
		return nil, err
	}
	log.Info("SSH portfolio server created", "addr", addr)
	return srv, nil
}

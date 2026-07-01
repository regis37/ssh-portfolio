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

	mw := []wish.Middleware{
		bm.Middleware(ui.TeaHandler),
		activeterm.Middleware(),
		logging.Middleware(),
	}

	// Visitor logging is opt-in: both PORTFOLIO_LOG_PATH and
	// PORTFOLIO_LOG_SALT must be set. The salt is generated once on the
	// server and stored in /opt/portfolio/.env (never in git).
	logPath := os.Getenv("PORTFOLIO_LOG_PATH")
	logSalt := os.Getenv("PORTFOLIO_LOG_SALT")
	if logPath != "" && logSalt != "" {
		vl, err := openVisitorLog(logPath, logSalt)
		if err != nil {
			log.Warn("Visitor logging disabled", "error", err)
		} else {
			log.Info("Visitor logging enabled", "path", logPath)
			mw = append(mw, visitorMiddleware(vl)) // outermost — runs first
		}
	} else {
		log.Info("Visitor logging disabled (set PORTFOLIO_LOG_PATH + PORTFOLIO_LOG_SALT to enable)")
	}

	opts := []ssh.Option{
		wish.WithAddress(addr),
		wish.WithMiddleware(mw...),
	}

	if hk := os.Getenv("PORTFOLIO_HOST_KEY"); hk != "" {
		opts = append(opts, wish.WithHostKeyPath(hk))
	}

	srv, err := wish.NewServer(opts...)
	if err != nil {
		return nil, err
	}
	log.Info("SSH portfolio server created", "addr", addr)
	return srv, nil
}

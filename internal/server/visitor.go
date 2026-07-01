package server

import (
	"crypto/sha256"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
)

// visitorLog writes one log line per SSH connection.
//
// Privacy / GDPR-DSGVO: IP addresses are personal data under Art. 4(1) GDPR.
// We never store them in plaintext. Instead we compute SHA-256(salt || ip),
// where the salt is a random 32-byte secret kept only on the server
// (env var PORTFOLIO_LOG_SALT, never committed to git).
// This satisfies GDPR recital 26: the hash is not "reasonably linkable" to a
// natural person without the salt, so it is treated as pseudonymous data.
// We truncate to the first 16 hex chars (64-bit prefix) — enough collision
// resistance for unique-visitor counting at this traffic scale.
type visitorLog struct {
	mu   sync.Mutex
	f    *os.File
	salt string
}

func openVisitorLog(path, salt string) (*visitorLog, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, fmt.Errorf("mkdir %s: %w", filepath.Dir(path), err)
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o640)
	if err != nil {
		return nil, fmt.Errorf("open log: %w", err)
	}
	return &visitorLog{f: f, salt: salt}, nil
}

func (vl *visitorLog) record(addr net.Addr) {
	host, _, err := net.SplitHostPort(addr.String())
	if err != nil {
		host = addr.String()
	}
	sum := sha256.Sum256([]byte(vl.salt + host))
	hash := fmt.Sprintf("%x", sum)[:16]

	line := time.Now().UTC().Format(time.RFC3339) + " " + hash + "\n"

	vl.mu.Lock()
	defer vl.mu.Unlock()
	_, _ = vl.f.WriteString(line)
}

// visitorMiddleware must be appended LAST to wish.WithMiddleware so it runs
// first in the handler chain, capturing every connection (including those
// later rejected by activeterm for missing PTY).
func visitorMiddleware(vl *visitorLog) wish.Middleware {
	return func(next ssh.Handler) ssh.Handler {
		return func(sess ssh.Session) {
			vl.record(sess.RemoteAddr())
			next(sess)
		}
	}
}

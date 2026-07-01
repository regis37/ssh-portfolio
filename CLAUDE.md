# SSH Portfolio

A terminal portfolio accessible via SSH, built with Go + Charmbracelet stack.

## Stack
- **Go 1.22**
- **wish** — SSH server framework
- **bubbletea** — TUI framework
- **lipgloss** — terminal styling

## Structure
```
cmd/portfolio/main.go          — entrypoint
internal/server/server.go     — wish server config + middleware
internal/ui/
  model.go                    — app state
  update.go                   — keyboard event handling
  view.go                     — render layout
  styles.go                   — lipgloss styles (light/dark themes)
  content.go                  — portfolio section content
```

## Run
```bash
go run ./cmd/portfolio
ssh -p 23234 localhost
```

## Build
```bash
go build ./...
```

## Connect
```bash
ssh -p 23234 127.0.0.1
```

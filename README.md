# ssh-portfolio

> Terminal portfolio — connect via SSH

```
ssh -p 23234 127.0.0.1
```

## Controls

| Key       | Action              |
|-----------|---------------------|
| `j` / `↓` | Next section       |
| `k` / `↑` | Previous section   |
| `Enter`   | Open section       |
| `Tab`     | Toggle panel focus |
| `t`       | Toggle theme       |
| `?`       | Help               |
| `q`       | Quit               |

## Build & run

```bash
go mod tidy
go build ./...
./portfolio
```

// stats reads the visitor log and prints aggregate analytics.
// Not exposed in the public TUI — run directly on the server:
//
//	/opt/portfolio/stats [/var/log/portfolio/visits.log]
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const defaultLogPath = "/var/log/portfolio/visits.log"

func main() {
	path := defaultLogPath
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	var (
		total  int
		unique = make(map[string]struct{})
		perDay = make(map[string]int)
	)

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		ts, hash := fields[0], fields[1]
		day := ts
		if len(ts) >= 10 {
			day = ts[:10]
		}
		total++
		unique[hash] = struct{}{}
		perDay[day]++
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "scan error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Total connections : %d\n", total)
	fmt.Printf("Unique visitors   : %d  (by 64-bit IP hash — GDPR-safe)\n", len(unique))

	if len(perDay) == 0 {
		fmt.Println("No data.")
		return
	}

	fmt.Println()
	fmt.Println("Connections per day:")
	days := make([]string, 0, len(perDay))
	for d := range perDay {
		days = append(days, d)
	}
	sort.Strings(days)
	for _, d := range days {
		bar := strings.Repeat("█", min(perDay[d], 40))
		fmt.Printf("  %s  %-3d %s\n", d, perDay[d], bar)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

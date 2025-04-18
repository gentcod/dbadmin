package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"strings"

	"github.com/gentcod/db-admin/cmd/start"
	cmd "github.com/gentcod/db-admin/internal/cli/root"
	"github.com/gentcod/db-admin/internal/config"
)

func main() {
	dbDriver, dbConnString := start.Execute()
	conn, err := sql.Open(dbDriver, dbConnString)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to initialize db: %w", err)
	}

	defer conn.Close()
	if err := conn.Ping(); err != nil {
		fmt.Fprintln(os.Stderr, "unable to connect to db: %w", err)
	}

	config.Cfg = config.NewConfig(conn)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nInteractive DbAdmin CLI (type 'exit' to quit)")
	for {
		fmt.Print("\ndbadmin> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		switch input {
		case "exit":
			return
		}

		args := strings.Split(input, " ")
		os.Args = append([]string{os.Args[0]}, args...)
		cmd.Execute()
	}
}

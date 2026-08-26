package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/paugomez86/gator/internal/config"
	"github.com/paugomez86/gator/internal/database"
	"github.com/paugomez86/gator/internal/repl"
)

func main() {
	var state config.State
	var commands repl.Commands
	var command repl.Command
	var err error

	if len(os.Args) < 2 {
		fmt.Printf("Invalid argument\n")
		os.Exit(1)
	}
	command.Name = os.Args[1]
	command.Args = os.Args[2:]

	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	state.Cfg = &cfg

	// Opening database connection
	db, err := sql.Open("postgres", state.Cfg.DbUrl)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	state.Db = database.New(db)

	commands.Map = make(map[string]func(*config.State, repl.Command) error)
	commands.Register("login", repl.HandlerLogin)
	commands.Register("register", repl.HandlerRegister)
	commands.Register("reset", repl.HandlerReset)
	commands.Register("users", repl.HandlerUsers)

	if err = commands.Run(&state, command); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

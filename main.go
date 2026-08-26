package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/paugomez86/gator/internal/config"
	"github.com/paugomez86/gator/internal/database"
	"github.com/paugomez86/gator/internal/handlers"
)

func main() {
	var state config.State
	var commands handlers.Commands
	var command handlers.Command
	var err error

	// Handling arguments
	if len(os.Args) < 2 {
		fmt.Printf("Invalid argument\n")
		os.Exit(1)
	}
	command.Name = os.Args[1]
	command.Args = os.Args[2:]

	// Reading config file
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

	// Registering commands
	commands.Map = make(map[string]func(*config.State, handlers.Command) error)
	commands.Register("login", handlers.HandlerLogin)
	commands.Register("register", handlers.HandlerRegister)
	commands.Register("reset", handlers.HandlerReset)
	commands.Register("users", handlers.HandlerUsers)

	// Running given command
	if err = commands.Run(&state, command); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

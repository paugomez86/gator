package repl

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/paugomez86/gator/internal/config"
	"github.com/paugomez86/gator/internal/database"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Map map[string]func(*config.State, Command) error
}

func HandlerLogin(s *config.State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Invalid argument")
	}

	user, err := s.Db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("User not found")
		}
		return err
	}

	if err := s.Cfg.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Printf("User set: %v\n", cmd.Args[0])

	return nil
}

func HandlerRegister(s *config.State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Invalid argument")
	}

	var args database.CreateUserParams

	args.ID = uuid.New()
	args.CreatedAt = time.Now()
	args.UpdatedAt = time.Now()
	args.Name = cmd.Args[0]

	user, err := s.Db.CreateUser(context.Background(), args)
	if err != nil {
		return err
	}
	s.Cfg.SetUser(user.Name)
	fmt.Printf("User created:\nID: %v\nName: %v\nTime: %v\n", user.ID, user.Name, user.CreatedAt)
	return nil
}

func (c Commands) Run(s *config.State, cmd Command) error {
	if command, ok := c.Map[cmd.Name]; ok {
		if err := command(s, cmd); err != nil {
			return fmt.Errorf("%v", err)
		}

	} else {
		return fmt.Errorf("Invalid command")
	}
	return nil
}

func (c Commands) Register(name string, f func(*config.State, Command) error) {
	c.Map[name] = f
}

package core

import (
	"bufio"
	"fmt"
	"homework_4/internal/commands"
	"homework_4/internal/commands/gofmt"
	"homework_4/internal/commands/help"
	"homework_4/internal/commands/spell"
	"os"
	"strings"
)

type CLI struct {
	Commands map[string]commands.CommandInterface
}

func NewCLI() *CLI {
	cli := &CLI{
		Commands: make(map[string]commands.CommandInterface),
	}
	helpCmd := help.NewHelpCommand(cli.Commands)
	cli.AddCommand(helpCmd)
	cli.AddCommand(spell.NewSpell())
	cli.AddCommand(gofmt.NewGofmtCommand())
	return cli
}

func (c *CLI) AddCommand(cmd commands.CommandInterface) {
	c.Commands[cmd.GetName()] = cmd
}

func (c *CLI) Execute(commandName string, args []string) error {
	cmd, ok := c.Commands[commandName]
	if !ok {
		return fmt.Errorf("Unknown command: %s", commandName)
	}
	arguments := make(map[string]string)
	fmt.Println(args)
	for i, arg := range args {
		if !strings.HasPrefix(arg, "--") {
			continue
		}
		key := string(arg[2:])
		if key == "" {
			return fmt.Errorf("missing name of argument: %s", key)
		}
		_, ok := cmd.GetRequired()[key]
		if !ok {
			return fmt.Errorf("%s is invalid  argument for command %s", key, commandName)
		}
		if i+1 < len(args) {
			arguments[key] = args[i+1]
			continue
		}
		return fmt.Errorf("missing required value for argument: %s", key)
	}
	return cmd.DoAction(arguments)
}
func (c *CLI) Run() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the CLI application. Type 'help' for a list of available commands.")

	for {
		fmt.Print("-> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		args := strings.Fields(input)
		cmdName := args[0]

		if cmdName == "exit" {
			fmt.Println("Goodbye!")
			return
		}
		if err := c.Execute(cmdName, args[1:]); err != nil {
			fmt.Fprintln(os.Stderr, "Error executing command:", err)
		}
	}
}

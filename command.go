package main

import "fmt"

type Command interface {
	Execute() string
	RoolBack() string
}

type CreateCommand struct{}

func (p *CreateCommand) Execute() string {
	return "CreateCommand"
}
func (p *CreateCommand) RoolBack() string {
	return "RoolBack CreateCommand"
}

type BindCommand struct{}

func (p *BindCommand) Execute() string {
	return "BindCommand"
}
func (p *BindCommand) RoolBack() string {
	return "RoolBack BindCommand"
}

type UnBindCommand struct{}

func (p *UnBindCommand) Execute() string {
	return "UnBindCommand"
}
func (p *UnBindCommand) RoolBack() string {
	return "RoolBack UnBindCommand"
}

type DropCommand struct{}

func (p *DropCommand) Execute() string {
	return "DropCommand"
}

func (p *DropCommand) RoolBack() string {
	return "RoolBack DropCommand"
}

func execCommands() {
	// Register commands
	commands := [...]Command{
		&DropCommand{},
		&UnBindCommand{},
		&BindCommand{},
		&CreateCommand{},
	}

	for _, command := range commands {
		fmt.Printf("Execute : %s", command.Execute())
		fmt.Println()
	}

	for _, command := range commands {
		fmt.Printf("RoolBack : %s", command.RoolBack())
		fmt.Println()

	}

}

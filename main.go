package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config) error
}

type config struct {
	next string
	prev any
}

var cfg = config{
	next: "",
	prev: nil,
}

func setConfig(next string, prev any) {
	cfg.next = next
	cfg.prev = prev
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	printPrompt()
	for {
		reader.Scan()
		input := reader.Text()
		command, exists := getCommands()[input]

		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func printPrompt() {
	fmt.Print("> ")
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Gets the names of the next 20 location areas in the Pokemon world.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Gets the names of the previous 20 locations areas in the Pokemon world.",
			callback:    commandBMap,
		},
	}
}

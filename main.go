package main

import (
	"bufio"
	"fmt"
	"os"
)

func failOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	fmt.Println("Welcome to shark_sandwich!")

	ConsoleReader := bufio.NewReader(os.Stdin)
	storage, err := NewStorage()
	failOnError(err)

	hero, err := InitGame(ConsoleReader, storage)
	failOnError(err)

	fmt.Println("Here are your measurements")
	fmt.Printf("%+v\n", hero)
	fmt.Println("Reminder: You can type 'help' at any time to get a list of options.")
	commandHelp := new(CommandHelp)
	commandHelp.Init()
	commandHelp.PrintHelpCommands()

	// todo: repl loop to deal with commands
	fmt.Print("Please enter command: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "quit" || line == "q" {
			break
		}
		fmt.Print("Please enter command: ")
		// do something with the command in a switch statement
	}
}

func printCommands() {
	fmt.Println()
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	ConsoleReader := bufio.NewReader(os.Stdin)
	storage := NewStorage()

	// todo: should be prompted to load an existing hero here as well
	fmt.Print("Welcome to shark_sandwich! Looks like you're new. Tell us about your hero so you can get started. What's your name? ")
	heroName, err := ConsoleReader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	hero := NewHero(heroName)
	storage.StorePlayer(*hero)

	fmt.Println("That's it! You're ready to go on an adventure.")
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
		fmt.Print("Please enter command: ")
		line := scanner.Text()
		if line == "quit" {
			break
		}
		// do something with the command in a switch statement
	}
}

func printCommands() {
	fmt.Println()
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	ConsoleReader := bufio.NewReader(os.Stdin)

	// todo: should be prompted to load an existing hero here as well
	fmt.Println("Welcome to shark_sandwich! Looks like you're new. Tell us about your hero so you can get started. What's your name?")
	heroName, err := ConsoleReader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	hero := NewHero(heroName)

	fmt.Println("That's it! You're ready to go on an adventure.")
	fmt.Println("Here are your measurements")
	fmt.Printf("%+v\n", hero)

	// todo: repl loop to deal with commands
	//command, err := ConsoleReader.ReadString('\n')
	command := ""
	for command != "quit" {
		fmt.Println("Please enter command: ")
		command, err := ConsoleReader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		// debug
		fmt.Println(command)
	}
}

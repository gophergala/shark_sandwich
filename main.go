package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	dir, err := os.Getwd()
	failOnError(err)

	err = storage.OpenRepository(dir)
	if err != nil {
		loadGame(ConsoleReader, storage)
	}

	// todo: should be prompted to load an existing hero here as well
	fmt.Print("Looks like you're new. Tell us about your hero so you can get started. What's your name? ")
	heroName, err := ConsoleReader.ReadString('\n')
	failOnError(err)

	heroName = strings.TrimSpace(heroName)
	hero := NewHero(heroName)
	err = storage.StorePlayer(*hero)
	failOnError(err)

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

func loadGame(ConsoleReader *bufio.Reader, storage *Storage) {
	fmt.Print("There is not a current game in this folder. Please enter a folder location to play the game in: ")
	folderPath, err := ConsoleReader.ReadString('\n')
	failOnError(err)

	folderPath = strings.TrimSpace(folderPath)
	err = storage.OpenRepository(folderPath)
	if err != nil {
		fmt.Print("There is not a current game in this folder. Please enter a remote url to load a game: ")
		remoteUrl, err := ConsoleReader.ReadString('\n')
		failOnError(err)

		remoteUrl = strings.TrimSpace(remoteUrl)
		err = storage.CloneRepository(remoteUrl, folderPath)
		failOnError(err)
	}
}

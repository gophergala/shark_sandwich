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
	fmt.Println()
	fmt.Println("Welcome to shark_sandwich!")
	fmt.Println()

	ConsoleReader := bufio.NewReader(os.Stdin)
	storage, err := NewStorage()
	failOnError(err)

	hero, err := InitGame(ConsoleReader, storage)
	failOnError(err)

	gameWorld := NewGameWorld(hero)

	gameLog := &GameLog{}
	gameLog.InitLogEventStream(gameWorld.SendLog)
	gameLog.PrintGameLog()
	fmt.Println("My Hero")
	fmt.Println("-------")
	fmt.Print(hero.String())
	commandHelp := new(CommandHelp)
	commandHelp.Init()
	commandHelp.PrintHelpCommands()

	pveFight := NewPveFight()
	gameWorld.addChannel(pveFight.SendEvent)
		
	// REPL
	fmt.Print("Please enter command: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text()

		switch command {
		case "adventure":
			a := NewAdventure(hero)
			a.Embark(pveFight)
			// todo: call adventure code and pass in channel to recieve game engine messages
			// todo: don't allow user to enter new command until adventure outcome is done (wait on event?)
			fmt.Print("Please enter command: ")
		case "help":
			commandHelp.PrintHelpCommands()
			fmt.Print("Please enter command: ")
		case "me":
			fmt.Println("My Hero")
			fmt.Print(hero.String())
			fmt.Println()
			fmt.Print("Please enter command: ")
		case "log":
			gameLog.PrintGameLog()
			commandHelp.PrintHelpCommands()
		case "quit", "q":
			fmt.Println("leaving so soon?")
			// todo: save game state
			os.Exit(0)
		default:
			fmt.Println("unknown command")
			fmt.Print("Please enter command: ")
		}
	}
}

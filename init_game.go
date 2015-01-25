package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func InitGame(ConsoleReader *bufio.Reader, storage *Storage) (*HeroSheet, error) {
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

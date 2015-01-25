package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func InitGame(ConsoleReader *bufio.Reader, storage *Storage) (*HeroSheet, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	err = storage.OpenRepository(dir)
	if err != nil {
		err = loadGame(ConsoleReader, storage)
		if err != nil {
			return nil, err
		}
	}

	_, err = storage.GetGameObject("shark_sandwich_game_data")
	if err != nil {
		err = loadGame(ConsoleReader, storage)
		if err != nil {
			return nil, err
		}
	}

	playerId, err := storage.GetCurrentPlayer()
	if err != nil {
		fmt.Print("Looks like you're new. Tell us about your hero so you can get started. What's your name? ")
		heroName, err := ConsoleReader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		playerId = strings.TrimSpace(heroName)
		hero := NewHero(playerId)
		err = storage.StorePlayer(*hero)
		if err != nil {
			return nil, err
		}

		err = storage.SetCurrentPlayer(playerId)
		if err != nil {
			return nil, err
		}

		fmt.Println("That's it! You're ready to go on an adventure.")
		if err != nil {
			return nil, err
		}
	}

	hero, err := storage.GetPlayer(playerId)
	if err != nil {
		return nil, err
	}

	return &hero, nil
}

func loadGame(ConsoleReader *bufio.Reader, storage *Storage) error {
	fmt.Print("There is not a current game in this folder. Please enter a folder location to play the game in: ")
	folderPath, err := ConsoleReader.ReadString('\n')
	if err != nil {
		return err
	}

	folderPath = strings.TrimSpace(folderPath)
	err = storage.OpenRepository(folderPath)
	if err != nil {
		fmt.Print("There is not a current game in this folder. Please enter a remote url to load a game: ")
		remoteUrl, err := ConsoleReader.ReadString('\n')
		if err != nil {
			return err
		}

		remoteUrl = strings.TrimSpace(remoteUrl)
		err = storage.CloneRepository(remoteUrl, folderPath)
		if err != nil {
			return err
		}
	}

	return nil
}

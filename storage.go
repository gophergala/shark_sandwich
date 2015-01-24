package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Storage struct {
	Events <-chan string
}

func NewStorage() *Storage {
	storage := &Storage{
		Events: make(chan string),
	}
	storage.initEventStream()
	return storage
}

func (s *Storage) initEventStream() {
	go func() {
		for _ = range s.Events {

		}
	}()
}

func (s *Storage) Clone(repoUrl string) error {
	// todo: Clone repo with provided url
	return nil
}

func (s *Storage) StorePlayer(hero HeroSheet) error {
	file, err := os.Create("players/" + hero.Name + "/" + hero.Name)
	if err != nil {
		return err
	}
	defer file.Close()

	heroBytes, err := json.Marshal(hero)
	if err != nil {
		return err
	}

	_, err = file.WriteString(string(heroBytes))
	if err != nil {
		return err
	}

	file.Sync()
	return nil

	// todo: Git commit and git push
}

func (s *Storage) StoreEvent(hero HeroSheet, event string) error {
	// todo: Git commit new event on the player file. git push to remote
	return nil
}

func (s *Storage) GetNewUpdates() error {
	// todo: git fetch
	return nil
}

func (s *Storage) GetPlayer(playerId string) (HeroSheet, error) {
	return getPlayer(playerId)
}

func (s *Storage) GetGameObject(id string) (string, error) {
	// todo: load and return file from file system
	return "", nil
}

func (s *Storage) GetGameObjects(id []string) ([]string, error) {
	// todo: load and return multiple file from file system
	return make([]string, 0), nil
}

func getPlayer(playerId string) (HeroSheet, error) {
	contents, _ := ioutil.ReadFile("players/" + playerId + "/" + playerId)
	heroSheet := HeroSheet{}
	err := json.Unmarshal(contents, &heroSheet)

	if err != nil {
		return HeroSheet{}, err
	}

	return heroSheet, nil
}

package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Event struct {
	PlayerId string
	Message  string
}

type Storage struct {
	Events <-chan Event
}

func NewStorage() *Storage {
	storage := &Storage{
		Events: make(chan Event),
	}
	storage.initEventStream()
	return storage
}

func (s *Storage) initEventStream() {
	go func() {
		for event := range s.Events {
			s.storeEvent(event)
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

func (s *Storage) storeEvent(event Event) error {
	filename := "players/" + event.PlayerId + "/" + event.PlayerId + "events"
	file := &os.File{}
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		file, err = os.Create(filename)
		if err != nil {
			return err
		}
	} else {
		file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			return err
		}
	}
	defer file.Close()

	file.WriteString(event.Message + "\n")

	// todo: Git commit and git push to remote
	return nil
}

func (s *Storage) GetNewUpdates() error {
	// todo: git fetch
	return nil
}

func (s *Storage) GetPlayer(playerId string) (HeroSheet, error) {
	contents, err := s.getFileContents("players/" + playerId + "/" + playerId)
	if err != nil {
		return HeroSheet{}, err
	}

	heroSheet := HeroSheet{}
	err = json.Unmarshal(contents, &heroSheet)
	if err != nil {
		return HeroSheet{}, err
	}

	return heroSheet, nil
}

func (s *Storage) GetGameObject(id string) ([]byte, error) {
	contents, err := s.getFileContents(id)
	if err != nil {
		return nil, err
	}

	return contents, nil
}

func (s *Storage) getFileContents(filename string) ([]byte, error) {
	contents, err := ioutil.ReadFile(filename)
	return contents, err
}

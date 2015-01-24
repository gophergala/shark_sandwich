package main

import (
	"encoding/json"
	"os"
)

type Storage struct{}

func (s *Storage) Clone(repoUrl string) error {
	// todo: Clone repo with provided url
}

func (s *Storage) StorePlayer(player Player) error {
	f, err := os.Create("players/" + player.Name)
	check(err)
	defer f.Close()
	_, err := f.WriteString(json.Marshall(player))

	if err != nil {
		fmt.Println("error:", err)
		return err
	}

	f.Sync()

	// todo: Git commit and git push
}

func (s *Storage) StoreEvent(player Player, event PlayerEvent) error {
	// todo: Git commit new event on the player file. git push to remote
}

func (s *Storage) GetNewUpdates() error {
	// todo: git fetch
}

func (s *Storage) GetPlayer(playerId string) (Player, error) {
	// todo: load and return a player from file system
	
	contents,_ := ioutil.ReadFile("players/" + playerId)
	var player Player()
	err := json.Unmarshal(contents, &player)
	
	if err != nil {
		return err
	}
	
	return player
}

func (s *Storage) GetPlayers(player []string) ([]Player, error) {
	// todo: load and return multiple players from file system
}

func (s *Storage) GetGameObject(id string) (string, error) {
	// todo: load and return file from file system
}

func (s *Storage) GetGameObjects(id []string) ([]string, error) {
	// todo: load and return multiple file from file system
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

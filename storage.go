package main

type Storage struct {}

func (s *Storage) Clone(repoUrl string) error {
	// todo: Clone repo with provided url
}

func (s *Storage) StorePlayer(player Player) error {
	// todo: Git commit and git push
}

func (s *Storage) StoreEvent(player Player, event PlayerEvent) error {
	// todo: Git commit new event on the player file. git push to remote
}

func (s *Storage) GetNewUpdates() error {
	// todo: git fetch
}

func (s *Storage) GetPlayer(player string) (Player, error) {
	// todo: load and return a player from file system
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

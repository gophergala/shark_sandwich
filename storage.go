package main

import (
	"encoding/json"
	"github.com/libgit2/git2go"
	"io/ioutil"
	"os"
)

const CURRENT_PLAYER_CONFIG_KEY string = "current_game_player"

type Event struct {
	PlayerId string
	Message  string
}

type Storage struct {
	Events     <-chan Event
	repository *git.Repository
	config     *git.Config
}

func NewStorage() (*Storage, error) {
	configPath, err := git.ConfigFindGlobal()
	if err != nil {
		return nil, err
	}
	config, err := git.OpenOndisk(nil, configPath)
	if err != nil {
		return nil, err
	}

	storage := &Storage{
		Events: make(chan Event),
		config: config,
	}
	storage.initEventStream()

	return storage, nil
}

func (s *Storage) initEventStream() {
	go func() {
		for event := range s.Events {
			s.storeEvent(event)
		}
	}()
}

func (s *Storage) GetCurrentPlayer() (string, error) {
	val, err := s.config.LookupString(CURRENT_PLAYER_CONFIG_KEY)
	if err != nil {
		return "", err
	}

	return val, nil
}

func (s *Storage) SetCurrentPlayer(playerId string) error {
	if err := s.config.SetString(CURRENT_PLAYER_CONFIG_KEY, playerId); err != nil {
		return err
	}

	return nil
}

func (s *Storage) OpenRepository(path string) error {
	repo, err := git.OpenRepository(path)
	if err != nil {
		return err
	}
	s.repository = repo

	return nil
}

func (s *Storage) CloneRepository(repoUrl string, path string) error {
	checkoutOptions := &git.CheckoutOpts{
		Strategy: git.CheckoutSafe,
	}
	cloneOptions := &git.CloneOptions{
		Bare:           false,
		CheckoutBranch: "master",
		CheckoutOpts:   checkoutOptions,
	}

	repo, err := git.Clone(repoUrl, path, cloneOptions)
	if err != nil {
		return err
	}
	repo.CreateRemote("origin", path)
	s.repository = repo

	return nil
}

func (s *Storage) NewRepository(path string) error {
	repo, err := git.InitRepository(path, false)
	if err != nil {
		return err
	}
	s.repository = repo

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

	err = s.commitCurrentIndex()
	if err != nil {
		return err
	}

	return s.pushLatestCommits()
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

	err = s.commitCurrentIndex()
	if err != nil {
		return err
	}

	return s.pushLatestCommits()
}

func (s *Storage) GetNewUpdates() error {
	remote, err := s.repository.LookupRemote("origin")
	if err != nil {
		return err
	}

	return remote.Fetch([]string{"master"}, nil, "")
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

func (s *Storage) commitCurrentIndex() error {
	index, err := s.repository.Index()
	if err != nil {
		return err
	}

	err = index.AddAll([]string{}, git.IndexAddDefault, nil)
	if err != nil {
		return err
	}

	_, err = index.WriteTreeTo(s.repository)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) pushLatestCommits() error {
	remote, err := s.repository.LookupRemote("origin")
	if err != nil {
		return err
	}

	return remote.Push([]string{"refs/heads/master"}, nil, nil, "")
}

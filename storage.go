package main

import (
	"encoding/json"
	"github.com/libgit2/git2go"
	"io/ioutil"
	"os"
	"time"
)

const CURRENT_PLAYER_CONFIG_KEY string = "current.game.player"

type Event struct {
	PlayerId string
	Message  string
}

type Storage struct {
	Events     <-chan Event
	repository *git.Repository
	path       string
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
	s.path = path

	return nil
}

func (s *Storage) CloneRepository(repoUrl string, path string) error {
	checkoutOptions := &git.CheckoutOpts{
		Strategy: git.CheckoutSafeCreate,
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
	s.path = path

	return nil
}

func (s *Storage) StorePlayer(hero HeroSheet) error {
	filepath := s.path + "/players/" + hero.Name
	err := os.MkdirAll(filepath, 0755)
	if err != nil {
		return err
	}

	filename := filepath + "/" + hero.Name
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	heroBytes, err := json.Marshal(hero)
	if err != nil {
		return err
	}

	_, err = file.WriteString(string(heroBytes) + "\n")
	if err != nil {
		return err
	}

	file.Sync()
	err = s.commitCurrentIndex("Added new player: " + hero.Name)
	if err != nil {
		return err
	}

	return s.pushLatestCommits()
}

func (s *Storage) storeEvent(event Event) error {
	err := os.MkdirAll(s.path+"/players/"+event.PlayerId, 0755)
	if err != nil {
		return err
	}

	filename := s.path + "/players/" + event.PlayerId + "/" + event.PlayerId + "events"
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
	err = s.commitCurrentIndex("Event: " + event.Message)
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

func (s *Storage) GetPlayer(playerId string) (*HeroSheet, error) {
	contents, err := s.getFileContents(s.path + "/players/" + playerId + "/" + playerId)
	if err != nil {
		return nil, err
	}

	heroSheet := HeroSheet{}
	err = json.Unmarshal(contents, &heroSheet)
	if err != nil {
		return nil, err
	}

	return &heroSheet, nil
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

func (s *Storage) commitCurrentIndex(message string) error {
	signature := &git.Signature{
		Name:  "shark_sandwich_engine",
		Email: "shark@sandwich.com",
		When:  time.Now(),
	}

	index, err := s.repository.Index()
	if err != nil {
		return err
	}

	err = index.AddAll([]string{}, git.IndexAddDefault, nil)
	if err != nil {
		return err
	}

	treeId, err := index.WriteTree()
	if err != nil {
		return err
	}

	err = index.Write()
	if err != nil {
		return err
	}

	tree, err := s.repository.LookupTree(treeId)
	if err != nil {
		return err
	}

	head, err := s.repository.Head()
	if err != nil {
		return err
	}

	commitTarget, err := s.repository.LookupCommit(head.Target())
	if err != nil {
		return err
	}

	_, err = s.repository.CreateCommit("HEAD", signature, signature, message, tree, commitTarget)
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

	//return remote.Push([]string{"refs/heads/master"}, nil, nil, "")
	return nil
}

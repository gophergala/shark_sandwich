package main

import "fmt"

type (
	FightEvent struct {
		Won bool
		Event
	}

	GameWorld struct {
		Hero *HeroSheet
		SendEvent chan string
	}
)

func NewGameWorld(hero *HeroSheet) *GameWorld {
	return &GameWorld{hero, make(chan string, 10)}
}

func (g *GameWorld) addChannel(events chan interface{}) {
	go func() {
		for {
			e := <- events
			switch event := e.(type) {
			case FightEvent:
				fmt.Printf("Fight: %s\n", event.String())
				g.SendEvent <- event.String()
				if event.Won {
					g.Hero.Xp = g.Hero.Xp + 10
				}
			}
		}
	}()
}

func (f *FightEvent) String() string {
	if f.Won {
		return "You Won"
	} else {
		return "You Lost"
	}
}

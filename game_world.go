package main

import "fmt"

type (
	FightEvent struct {
		Won bool
		Event
	}
	
)

func GameWorld(hero *HeroSheet) func(interface{}) {
	// tried to do this async with channels but not working.
	// good enough for now
	return func(e interface{}) {
		switch event := e.(type) {
		case FightEvent:
			fmt.Printf("Fight: %s\n", event.String())
			hero.Xp = hero.Xp + 10
		}
	}
}

func (f *FightEvent) String() string {
	if f.Won {
		return "You Won"
	} else {
		return "You Lost"
	}
}

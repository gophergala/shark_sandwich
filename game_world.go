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
		SendLog chan LogEvent
	}
)

func NewGameWorld(hero *HeroSheet) *GameWorld {
	return &GameWorld{hero, make(chan string, 10), make(chan LogEvent, 100)}
}

func (g *GameWorld) addChannel(events chan interface{}) {
	go func() {
		for {
			e := <- events
			switch event := e.(type) {
			case FightEvent:
				message := fmt.Sprintf("Fight: %s\n", event.String())
				g.SendEvent <- event.String()
				if event.Won {
					g.Hero.Xp = g.Hero.Xp + 10
				}
				log := LogEvent {
					message,
					int(g.Hero.Xp),
					int(g.Hero.Life),
					int(g.Hero.Speed),
					int(g.Hero.Power),
					int(g.Hero.Ancestry),
				}
				g.SendLog <- log
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

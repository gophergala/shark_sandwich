package main

import (
	"fmt"
)

const (
	ADVENTURE_TYPE_ENCOUNTER AdventureType = 1
	ADVENTURE_TYPE_WANDER    AdventureType = 2
	ADVENTURE_TYPE_DISCOVERY AdventureType = 3
)

type (
	AdventureType int
	Adventure     struct {
		Type AdventureType
		*HeroSheet
	}
)

func NewAdventure(h *HeroSheet) *Adventure {
	return &Adventure{
		Type:      generateAdventure(),
		HeroSheet: h,
	}
}

func generateAdventure() AdventureType {
	return AdventureType(random(1, 3))
}

func (a *Adventure) Embark(pve *PveFight) {
	switch a.Type {
	case ADVENTURE_TYPE_DISCOVERY:
		fmt.Println("You didn't discover anything, too bad.")
	case ADVENTURE_TYPE_ENCOUNTER:
		fmt.Println("A wild Snorlax appeared.")
		snorlax := NewEnemy(a.HeroSheet)
		if snorlax.Life > a.HeroSheet.BaseStats.Life {
			fmt.Println(" He's a tough one!")
		}
		if snorlax.Power > a.HeroSheet.BaseStats.Power {
			fmt.Println(" He hits pretty hard...")
		}
		if snorlax.Speed > a.HeroSheet.BaseStats.Speed {
			fmt.Println(" Faster than your average snorlax.")
		}
		fmt.Println("Snorlax attacks you!")
		pve.Fight(a.HeroSheet, snorlax)
	case ADVENTURE_TYPE_WANDER:
		fmt.Println("You wandered right back to where you started")
	}
}

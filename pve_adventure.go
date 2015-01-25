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
		Type: generateAdventure(),
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
		fmt.Println("A wild Snorlax appeared and you fought!")
		snorlax := NewEnemy(a.HeroSheet)
		pve.Fight(a.HeroSheet, snorlax)
	case ADVENTURE_TYPE_WANDER:
		fmt.Println("You wandered right back to where you started")
	}
}

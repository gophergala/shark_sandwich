package main

import ()

const (
	ADVENTURE_TYPE_ENCOUNTER AdventureType = 1
	ADVENTURE_TYPE_WANDER    AdventureType = 2
	ADVENTURE_TYPE_DISCOVERY AdventureType = 3
)

type (
	AdventureType int
	Adventure     struct {
		Type AdventureType
	}
)

func NewAdventure(h *HeroSheet) *Adventure {
	return &Adventure{
		Type: generateAdventure(),
	}
}

func (a *Adventure) generateAdventure() (error, AdventureType) {
	return AdventureType(random(1, 3))
}

func (a *Adventure) Embark() {
	switch a.Type {
	case ADVENTURE_TYPE_DISCOVERY:
	case ADVENTURE_TYPE_ENCOUNTER:
	case ADVENTURE_TYPE_WANDER:

	}
}

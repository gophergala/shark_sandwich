package main

import ()

const (
	ADVENTURE_OUTCOME_ENCOUNTER AdventureOutcome = 1
	ADVENTURE_OUTCOME_WANDER    AdventureOutcome = 2
	ADVENTURE_OUTCOME_DISCOVERY AdventureOutcome = 3
)

type (
	AdventureOutcome int
	Adventure        struct {
	}
)

func NewAdventure(h *HeroSheet) *Adventure {
	return &Adventure{}
}

func (a *Adventure) Embark() (error, AdventureOutcome) {
	outcome := AdventureOutcome(random(1, 4))

	return nil, outcome
}

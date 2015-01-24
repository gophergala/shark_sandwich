package main

import (
	"testing"
)

func TestAdventure(t *testing.T) {
	h := NewHero("Durdle")
	a := NewAdventure(h)

	if a == nil {
		t.Log("No adventure? Lame.")
		t.Fail()
	}
}

func TestAdventureHasOutcome(t *testing.T) {
	h := NewHero("Durdle")
	a := NewAdventure(h)

	err, outcome := a.Embark()

	if err != nil {
		t.Log("Error on adventure: %s", err)
		t.Fail()
	}

	switch outcome {
	case ADVENTURE_OUTCOME_DISCOVERY:
	case ADVENTURE_OUTCOME_ENCOUNTER:
	case ADVENTURE_OUTCOME_WANDER:
	default:
		t.Logf("Unsupported adventure outcome: %v", outcome)
		t.Fail()
	}
}

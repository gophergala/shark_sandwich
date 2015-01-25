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

func TestGenerateAdventureHasOutcome(t *testing.T) {
	a := generateAdventure()

	switch a {
	case ADVENTURE_TYPE_DISCOVERY:
	case ADVENTURE_TYPE_ENCOUNTER:
	case ADVENTURE_TYPE_WANDER:
	default:
		t.Logf("Unsupported adventure generated: %v", a)
		t.Fail()
	}
}

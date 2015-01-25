package main

import (
	"testing"
)

func TestFightIncreasesXp(t *testing.T) {
	h := NewHero("Durdle")
	e := NewEnemy(h)

	e.Life = 1

	world := GameWorld(h)
	
	Fight(h, e, world)

	if h.Xp != 10 {
		t.Log("Player did not gain XP")
		t.Fail()
	}
}

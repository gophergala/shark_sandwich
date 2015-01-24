package main

import (
	"testing"
)

func TestGenerateEnemy(t *testing.T) {
	h := NewHero("Durdle")
	e := NewEnemy(h)

	if e == nil {
		t.Log("Enemey was nil")
		t.Fail()
	}
}

func TestEnemyGenStats(t *testing.T) {
	h := NewHero("Durdle")
	e := NewEnemy(h)
	testLife(&e.BaseStats, t, h.BaseStats.Life-ENEMY_VARIANCE_LIFE, h.BaseStats.Life+ENEMY_VARIANCE_LIFE)
	testPower(&e.BaseStats, t, h.BaseStats.Power-ENEMY_VARIANCE_POWER, h.BaseStats.Power+ENEMY_VARIANCE_POWER)
	testSpeed(&e.BaseStats, t, h.BaseStats.Speed-ENEMY_VARIANCE_SPEED, h.BaseStats.Speed+ENEMY_VARIANCE_SPEED)
}

func Fight(p1 *HeroSheet, n *NPCUnit) {
	// p1 is fighting n
	// copy base stats to temp
	// loop and run combat formula

	// Speed, Power, Life
	// Loop to increment "initiative" by speed, when > hit with power

	p1Life := p1.Life
	nLife := n.Life

	// Synchronize the following goroutines with channels
	// select { } with channels

	// stop := make(chan bool)

	// "p1" sim hitting "n", which is just a dummy right now
	go func() {
		for nLife > 0 {
			for p1Initiative := int64(0); p1Initiative < 100; p1Initiative += p1.Speed {
				// hit!
				nLife -= p1.Power
			}
		}
	}()
	// "n" sim hitting "p1"
	go func() {
		for p1Life > 0 {
			for nInitiative := int64(0); nInitiative < 100; nInitiative += n.Speed {
				// hit!
				p1Life -= n.Power
			}
		}
	}()
}

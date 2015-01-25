package main

import (
	"fmt"
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

func TestFight(t *testing.T) {
	h := NewHero("Spinal Tap")
	e := NewEnemy(h)
	called := false

	callback := func(e interface{}) {
		called = true
	}

	Fight(h, e, callback)

	if called == false {
		t.Log("The fight did not complete")
		t.Fail()
	}
}

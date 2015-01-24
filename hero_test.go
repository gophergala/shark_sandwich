package main

import (
	"testing"
)

func TestHeroInitAncestry(t *testing.T) {
	hero := NewHero("Durdle")
	if hero.Ancestry != 1 {
		t.Logf("Hero did not start with ancestry 1, instead was: %v", hero.Ancestry)
		t.Fail()
	}
}

func TestHeroGenStatsInRanges(t *testing.T) {
	hero := NewHero("Durdle")
	testLife(hero, t)
	testPower(hero, t)
	testSpeed(hero, t)
}

func testLife(h *HeroSheet, t *testing.T) {
	if h.Life < 0 {
		t.Logf("Hero started with negative life: %v", h.Life)
		t.Fail()
	}

	if h.Life < HERO_MIN_LIFE {
		t.Logf("Hero started with less than %v life: %v", HERO_MIN_LIFE, h.Life)
		t.Fail()
	}

	if h.Life > HERO_MAX_LIFE {
		t.Logf("Hero started with more than %v life: %v", HERO_MAX_LIFE, h.Life)
		t.Fail()
	}
}

func testPower(h *HeroSheet, t *testing.T) {
	if h.Power < 0 {
		t.Logf("Hero started with negative power: %v", h.Power)
		t.Fail()
	}

	if h.Power < HERO_MIN_POWER {
		t.Logf("Hero started with less than %v power: %v", HERO_MIN_POWER, h.Power)
		t.Fail()
	}

	if h.Power > HERO_MAX_POWER {
		t.Logf("Hero started with more than %v power: %v", HERO_MAX_POWER, h.Power)
		t.Fail()
	}
}

func testSpeed(h *HeroSheet, t *testing.T) {
	if h.Speed < 0 {
		t.Logf("Hero started with negative speed: %v", h.Speed)
		t.Fail()
	}

	if h.Speed < HERO_MIN_SPEED {
		t.Logf("Hero started with less than %v speed: %v", HERO_MIN_SPEED, h.Speed)
		t.Fail()
	}

	if h.Speed > HERO_MAX_SPEED {
		t.Logf("Hero started with more than %v speed: %v", HERO_MAX_SPEED, h.Speed)
		t.Fail()
	}
}

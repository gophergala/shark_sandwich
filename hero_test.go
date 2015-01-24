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
	testLife(&hero.BaseStats, t, HERO_MIN_LIFE, HERO_MAX_LIFE)
	testPower(&hero.BaseStats, t, HERO_MIN_POWER, HERO_MAX_POWER)
	testSpeed(&hero.BaseStats, t, HERO_MIN_SPEED, HERO_MAX_SPEED)
}

func testLife(stats *BaseStats, t *testing.T, min, max int64) {
	if stats.Life < 0 {
		t.Logf("Stats started with negative life: %v", stats.Life)
		t.Fail()
	}

	if stats.Life < min {
		t.Logf("Stats started with less than %v life: %v", min, stats.Life)
		t.Fail()
	}

	if stats.Life > max {
		t.Logf("Stats started with more than %v life: %v", max, stats.Life)
		t.Fail()
	}
}

func testPower(stats *BaseStats, t *testing.T, min, max int64) {
	if stats.Power < 0 {
		t.Logf("Stats started with negative power: %v", stats.Power)
		t.Fail()
	}

	if stats.Power < min {
		t.Logf("Stats started with less than %v power: %v", min, stats.Power)
		t.Fail()
	}

	if stats.Power > max {
		t.Logf("Stats started with more than %v power: %v", max, stats.Power)
		t.Fail()
	}
}

func testSpeed(stats *BaseStats, t *testing.T, min, max int64) {
	if stats.Speed < 0 {
		t.Logf("Stats started with negative speed: %v", stats.Speed)
		t.Fail()
	}

	if stats.Speed < min {
		t.Logf("Stats started with less than %v speed: %v", min, stats.Speed)
		t.Fail()
	}

	if stats.Speed > max {
		t.Logf("Stats started with more than %v speed: %v", max, stats.Speed)
		t.Fail()
	}
}

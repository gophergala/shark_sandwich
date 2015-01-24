package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	HERO_MAX_LIFE  int64 = 100
	HERO_MIN_LIFE  int64 = 75
	HERO_MIN_POWER int64 = 8
	HERO_MAX_POWER int64 = 15
	HERO_MIN_SPEED int64 = 8
	HERO_MAX_SPEED int64 = 15
)

func init() {
	// this is called one time when the package initializes
	// seed once so we can get different random numbers
	rand.Seed(time.Now().Unix())
}

type (
	BaseStats struct {
		Name  string
		Life  int64
		Speed int64
		Power int64
	}
	HeroSheet struct {
		Ancestry int64
		BaseStats
	}
)

func NewHero(name string) *HeroSheet {
	hero := &HeroSheet{
		BaseStats: BaseStats{
			Name: name,
		},
		Ancestry: 1,
	}
	hero.genStats()
	return hero
}

func (h *HeroSheet) genStats() {
	h.BaseStats.Life = random(HERO_MIN_LIFE, HERO_MAX_LIFE)
	h.BaseStats.Power = random(HERO_MIN_POWER, HERO_MAX_POWER)
	h.BaseStats.Speed = random(HERO_MIN_SPEED, HERO_MAX_SPEED)
}

func random(min, max int64) int64 {
	val := rand.Int63n(max-min) + min
	return val
}

func (h *HeroSheet) String() string {
	s := fmt.Sprintf("\n\tName: %v\n\tLife: %v\n\tPower: %v\n\tSpeed: %v\n",
		h.BaseStats.Name,
		h.BaseStats.Life,
		h.BaseStats.Power,
		h.BaseStats.Speed)
	return s
}

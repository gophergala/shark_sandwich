package main

import ()

const (
	ENEMY_VARIANCE_LIFE  int64 = 50
	ENEMY_VARIANCE_POWER int64 = 5
	ENEMY_VARIANCE_SPEED int64 = 5
)

type (
	NPCUnit struct {
		IsNPC bool
		BaseStats
	}
)

func NewEnemy(h *HeroSheet) *NPCUnit {
	npc := &NPCUnit{
		IsNPC: true,
	}

	npc.genNPCStats(h)
	return npc
}

func (n *NPCUnit) genNPCStats(h *HeroSheet) {
	n.Life = random(h.Life-ENEMY_VARIANCE_LIFE, h.Life+ENEMY_VARIANCE_LIFE)
	n.Power = random(h.Power-ENEMY_VARIANCE_POWER, h.Power+ENEMY_VARIANCE_POWER)
	n.Speed = random(h.Speed-ENEMY_VARIANCE_SPEED, h.Speed+ENEMY_VARIANCE_SPEED)
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

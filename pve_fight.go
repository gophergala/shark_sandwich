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

package lib

import "github.com/google/uuid"

type Player struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Game struct {
	Id     uuid.UUID `json:"id"`
	Player Player    `json:"player"`
	World  string    `json:"world"`
	Units  []Unit    `json:"units"`
}

func (g Game) Victory() bool {
	for _, unit := range g.Units {
		if unit.Self() && unit.Leader() && unit.Died {
			return false
		}
	}
	return true
}

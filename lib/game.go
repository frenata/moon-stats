package lib

import (
	"fmt"

	"github.com/google/uuid"
)

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

func (g Game) String() string {
	result := "won!"
	if g.Victory() {
		result = "lost!"
	}
	out := fmt.Sprintf("(%s) %s played on %s and %s",
		g.Id, g.Player.Name, g.World, result)
	out = out + fmt.Sprintf("\n\nPieces Played:\nid  | played | died\n--------------------\n")

	for _, unit := range g.Units {
		if unit.Self() {
			out = out + fmt.Sprintln(unit)
		}
	}

	return out
}

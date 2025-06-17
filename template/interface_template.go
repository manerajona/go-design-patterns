package template

import (
	"math/rand/v2"
)

type Game interface {
	Start()
	HaveWinner() bool
	TakeTurn()
	Winner() int
}

func GameTemplate(g Game) int {
	g.Start()
	for !g.HaveWinner() {
		g.TakeTurn()
	}
	return g.Winner()
}

type chess struct {
	turn, startingPlayer, currentPlayer int
}

func (c *chess) Start() {
	c.startingPlayer = rand.IntN(2)
	c.currentPlayer = c.startingPlayer
	c.turn = 0
}

func (c *chess) HaveWinner() bool {
	return c.turn == 12 // reached max number of turns
}

func (c *chess) TakeTurn() {
	c.turn++
	c.currentPlayer = 1 - c.currentPlayer
}

func (c *chess) Winner() int {
	return c.currentPlayer
}

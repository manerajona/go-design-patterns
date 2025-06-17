package template

import (
	"testing"
)

// deterministicChess allows deterministic start for testing
type deterministicChess struct {
	chess
	startWith int
}

func (c *deterministicChess) Start() {
	c.startingPlayer = c.startWith
	c.currentPlayer = c.startingPlayer
}

// Test that the game always takes 12 turns and alternates players correctly.
func TestGameTemplate_Chess(t *testing.T) {
	for startWith := 0; startWith <= 1; startWith++ {
		game := &deterministicChess{startWith: startWith}
		winner := GameTemplate(game)
		if game.turn != 12 {
			t.Errorf("Expected 12 turns, got %d", game.turn)
		}
		// The winner should be the player who made the last move
		expectedWinner := (startWith + 12) % 2
		if winner != expectedWinner {
			t.Errorf("Expected winner %d, got %d (startWith=%d)", expectedWinner, winner, startWith)
		}
	}
}

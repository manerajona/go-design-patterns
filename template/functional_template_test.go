package template

import (
	"testing"
)

// Example: A simple two-player game with 5 turns, player 0 starts.
func TestGameTemplateV2(t *testing.T) {
	var turns, currentPlayer, gameWinner int
	const maxTurns = 5

	start := func() {
		turns = 0
		currentPlayer = 0
		gameWinner = 1
	}

	takeTurn := func() {
		turns++
		// Alternate player each turn
		currentPlayer = 1 - currentPlayer
	}

	haveWinner := func() bool {
		return turns == maxTurns
	}

	winner := func() int {
		return currentPlayer
	}

	result := GameTemplateV2(start, takeTurn, haveWinner, winner)

	if turns != maxTurns {
		t.Errorf("expected %d turns, got %d", maxTurns, turns)
	}
	expectedWinner := (0 + maxTurns) % 2 // Starting player alternates each turn
	if result != expectedWinner {
		t.Errorf("expected winner %d, got %d", expectedWinner, result)
	}
	if gameWinner != result {
		t.Errorf("winner function and result mismatch: got %d and %d", gameWinner, result)
	}
}

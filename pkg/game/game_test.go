package game

import (
	"testing"
)

func TestStartGame(t *testing.T) {
	g := Game{}

	g.StartGame()

	count := 0
	for _, row := range g.Board {
		for _, cell := range row {
			if cell != 0 {
				count++
			}
		}
	}

	if count != 2 {
		t.Errorf("StartGame needs to add 2 bricks, got %d", count)
	}
}

func BenchmarkStartGame(b *testing.B) {
	for n := 0; n < b.N; n++ {
		g := Game{}

		g.StartGame()
	}
}

func TestAddBrick(t *testing.T) {
	g := Game{}

	g.AddBrick()

	var foundBrick = false
	for _, row := range g.Board {
		for _, cell := range row {
			if cell != 0 {
				foundBrick = true
			}
		}
	}

	if !foundBrick {
		t.Error("Could not find brick after AddBrick")
	}
}

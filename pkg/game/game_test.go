package game

import (
	"reflect"
	"testing"
)

func createNGames(n int) []Game {
	var games []Game

	for i := 0; i < n; i++ {
		games = append(games, Game{})
	}

	return games
}

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
	games := createNGames(b.N)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		games[n].StartGame()
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

func BenchmarkAddBrick(b *testing.B) {
	games := createNGames(b.N)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		games[n].AddBrick()
	}
}

func TestRotateBoardClockwise(t *testing.T) {
	g := Game{
		Board: [4][4]int{
			[4]int{2, 4, 0, 16},
			[4]int{2, 2, 0, 0},
			[4]int{0, 0, 0, 0},
			[4]int{0, 0, 0, 0},
		},
	}

	g.rotateBoard(false)

	rotatedBoard := [4][4]int{
		[4]int{0, 0, 2, 2},
		[4]int{0, 0, 2, 4},
		[4]int{0, 0, 0, 0},
		[4]int{0, 0, 0, 16},
	}

	if g.Board != rotatedBoard {
		t.Errorf("expected rotated board to be %+v, but found %+v", rotatedBoard, g.Board)
	}
}

func TestRotateBoardCounterClockwise(t *testing.T) {
	g := Game{
		Board: [4][4]int{
			[4]int{2, 4, 0, 16},
			[4]int{2, 2, 0, 0},
			[4]int{0, 0, 0, 0},
			[4]int{0, 0, 0, 0},
		},
	}

	g.rotateBoard(true)

	rotatedBoard := [4][4]int{
		[4]int{16, 0, 0, 0},
		[4]int{0, 0, 0, 0},
		[4]int{4, 2, 0, 0},
		[4]int{2, 2, 0, 0},
	}

	if g.Board != rotatedBoard {
		t.Errorf("expected rotated board to be %+v, but found %+v", rotatedBoard, g.Board)
	}
}

func BenchmarkRotateBoard(b *testing.B) {
	g := Game{}
	g.StartGame()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		g.rotateBoard(false)
	}
}

func TestSlideLeft(t *testing.T) {
	boards := [][4][4]int{
		[4][4]int{
			[4]int{16, 0, 16, 0},
			[4]int{1024, 0, 1024, 0},
			[4]int{4, 2, 0, 0},
			[4]int{2, 2, 0, 0},
		},
		[4][4]int{
			[4]int{16, 16, 0, 0},
			[4]int{128, 256, 4, 0},
			[4]int{4, 2, 0, 0},
			[4]int{2, 2, 0, 0},
		},
		[4][4]int{
			[4]int{16, 0, 0, 0},
			[4]int{0, 0, 0, 0},
			[4]int{4, 2, 0, 0},
			[4]int{2, 2, 0, 0},
		},
	}

	results := [][4][4]int{
		[4][4]int{
			[4]int{32, 0, 0, 0},
			[4]int{2048, 0, 0, 0},
			[4]int{4, 2, 0, 0},
			[4]int{4, 0, 0, 0},
		},
		[4][4]int{
			[4]int{32, 0, 0, 0},
			[4]int{128, 256, 4, 0},
			[4]int{4, 2, 0, 0},
			[4]int{4, 0, 0, 0},
		},
		[4][4]int{
			[4]int{16, 0, 0, 0},
			[4]int{0, 0, 0, 0},
			[4]int{4, 2, 0, 0},
			[4]int{4, 0, 0, 0},
		},
	}

	for i, board := range boards {
		g := Game{Board: board}

		g.slideLeft()

		if !reflect.DeepEqual(results[i], g.Board) {
			t.Errorf("expected slided board to be %+v, but found %+v", results[i], g.Board)
		}
	}
}

func BenchmarkSlideLeft(b *testing.B) {
	games := createNGames(b.N)
	for _, game := range games {
		game.Board = [4][4]int{
			[4]int{32, 0, 32, 64},
			[4]int{2048, 0, 64, 0},
			[4]int{4, 2, 256, 0},
			[4]int{4, 0, 128, 0},
		}
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		games[n].slideLeft()
	}
}

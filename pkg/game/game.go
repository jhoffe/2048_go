package game

import (
	"fmt"
	"math/rand"
	"time"
)

type Move int

const (
	Up Move = iota
	Down
	Left
	Right
)

func (m Move) String() string {
	return [...]string{"up", "down", "left", "right"}[m]
}

type Game struct {
	Board          [4][4]int
	Score          int64
	Done           bool
	PrintAfterMove bool
}

func (g *Game) AddBrick() *Game {
	emptyCells := g.FindEmptyCells()

	if len(emptyCells) <= 0 {
		g.Done = true
		return g
	}

	var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	emptyCell := emptyCells[rnd.Intn(len(emptyCells))]
	g.Board[emptyCell/len(g.Board)][emptyCell%len(g.Board)] = 2

	return g
}

func (g *Game) StartGame() *Game {
	g.AddBrick()
	g.AddBrick()

	return g
}

func (g *Game) rotateBoard(counterClockwise bool) *Game {
	var rotatedBoard = [4][4]int{}

	for i, row := range g.Board {
		rotatedBoard[i] = [4]int{}
		for j := range row {
			if counterClockwise {
				rotatedBoard[i][j] = g.Board[j][len(g.Board)-i-1]
			} else {
				rotatedBoard[i][j] = g.Board[len(g.Board)-j-1][i]
			}
		}
	}

	g.Board = rotatedBoard

	return g
}

func (g *Game) rotateBoardN(n int, counterClockwise bool) *Game {
	for i := n; i >= 0; i-- {
		g.rotateBoard(counterClockwise)
	}

	return g
}

func (g *Game) FindEmptyCells() []int {
	var emptyCells = []int{}
	for i, row := range g.Board {
		for j, cell := range row {
			if cell == 0 {
				emptyCells = append(emptyCells, i*len(g.Board)+j)
			}
		}
	}

	return emptyCells
}

func (g *Game) slideLeft() *Game {
	board := g.Board

	for y, row := range board {

		stopMerge := 0

		for j := 1; j < len(row); j++ {

			if row[j] != 0 {
				for x := j; x > stopMerge; x-- {
					if board[y][x-1] == 0 {
						board[y][x-1] = board[y][x]
						board[y][x] = 0
					} else if board[y][x-1] == board[y][x] {
						board[y][x-1] += board[y][x]
						g.Score += int64(board[y][x-1])
						board[y][x] = 0
						stopMerge = x
						break
					} else {
						break
					}
				}
			}
		}
	}

	g.Board = board

	return g
}

func (g *Game) Move(move Move) *Game {
	switch move {
	case Up:
		g.rotateBoard(true)
		g.slideLeft()
		g.rotateBoard(false)
	case Down:
		g.rotateBoard(false)
		g.slideLeft()
		g.rotateBoard(true)
	case Left:
		g.slideLeft()
	case Right:
		g.rotateBoardN(2, true)
		g.slideLeft()
		g.rotateBoardN(2, false)
	}

	g.AddBrick()

	if g.PrintAfterMove {
		g.PrintBoard()
	}

	return g
}

func (g *Game) PrintBoard() *Game {
	for _, row := range g.Board {
		fmt.Println(fmt.Sprintf("%v", row))
	}

	fmt.Printf("Score: %d\n\n", g.Score)

	return g
}

func (g *Game) Copy() Game {
	return *g
}

func (g *Game) GetHighestBrickValue() int {
	var highestBrick int

	for y, _ := range g.Board {
		for x, _ := range g.Board {
			if g.Board[y][x] > highestBrick {
				highestBrick = g.Board[y][x]
			}
		}
	}

	return highestBrick
}

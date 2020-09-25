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

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func (g *Game) AddBrick() *Game {
	emptyCells := g.findEmptyCells()

	if len(emptyCells) <= 0 {
		g.Done = true
		return g
	}

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

func (g *Game) findEmptyCells() []int {
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

// func getActionFromScores(scores map[Move]int64) Move {
// 	var action Move

// 	for move, score := range scores {
// 		if score > scores[action] {
// 			action = move
// 		}
// 	}

// 	return action
// }

// func main() {
// 	R := flag.Int("iters", 200, "the number of iterations the tree should do")
// 	N := flag.Int("n", 50, "the sample size")
// 	D := flag.Int("depth", 200, "the depth of the tree")

// 	file, err := os.Create("results.csv")
// 	if err != nil {
// 		panic(err)
// 	}

// 	writer := csv.NewWriter(file)
// 	defer func() {
// 		writer.Flush()
// 		file.Close()
// 	}()

// 	if err := writer.Write([]string{"N", "D", "R", "Score", "Highest"}); err != nil {
// 		panic(err)
// 	}

// 	var actions = [4]Move{Up, Down, Left, Right}

// 	for j := 0; j < *N; j++ {
// 		for r := 5; r <= *R; r += 5 {
// 			for d := 5; d <= *D; d += 5 {
// 				fmt.Println(fmt.Sprintf("beginning N = %d, R = %d, D = %d", j+1, r, d))
// 				game := Game{}

// 				for !game.Done {
// 					scores := map[Move]int64{}

// 					for _, init_action := range actions {
// 						func(action Move) {
// 							for i := 0; i < *R; i++ {
// 								cg := game.Copy()

// 								cg.Move(action)

// 								for m := 0; m < *D; m++ {
// 									cg.Move(actions[rnd.Intn(3)])

// 									if cg.Done {
// 										break
// 									}
// 								}

// 								// for !cg.Done {
// 								// 	cg.Move(actions[rnd.Intn(3)])
// 								// }

// 								scores[action] += cg.Score
// 							}
// 						}(init_action)
// 					}

// 					action := getActionFromScores(scores)

// 					// fmt.Println(action, scores)

// 					game.Move(action)
// 				}

// 				row := []string{strconv.Itoa(j + 1), strconv.Itoa(d), strconv.Itoa(r), strconv.Itoa(int(game.Score)), strconv.Itoa(game.GetHighestBrickValue())}

// 				if err := writer.Write(row); err != nil {
// 					panic(err)
// 				}
// 			}

// 		}
// 	}
// }

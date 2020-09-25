package montecarlo

import (
	"math/rand"
	"sync"
	"time"

	"github.com/jhoffe/2048_go/pkg/game"
)

func getActionFromScores(scores map[game.Move]int64) game.Move {
	var action game.Move

	for move, score := range scores {
		if score > scores[action] {
			action = move
		}
	}

	return action
}

func ScoreRewardFunction(g game.Game) int64 {
	return g.Score
}

func Run(R, D int, rf func(game.Game) int64) (game.Game, int64) {
	before := time.Now()
	var actions = [4]game.Move{game.Up, game.Down, game.Left, game.Right}
	g := game.Game{}

	var mu sync.Mutex
	var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

	for !g.Done {
		scores := map[game.Move]int64{}

		for _, init_action := range actions {
			for r := 0; r < R; r++ {
				cg := g.Copy()

				cg.Move(init_action)

				if D == 0 {
					for !cg.Done {
						mu.Lock()
						cg.Move(actions[rnd.Intn(3)])
						mu.Unlock()
					}
				} else {
					for m := 0; m < D; m++ {
						mu.Lock()
						cg.Move(actions[rnd.Intn(3)])
						mu.Unlock()

						if cg.Done {
							break
						}
					}

				}

				scores[init_action] += rf(cg)
			}
		}

		action := getActionFromScores(scores)

		g.Move(action)
	}

	return g, time.Now().Sub(before).Milliseconds()
}

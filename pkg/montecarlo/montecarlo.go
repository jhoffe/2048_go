package montecarlo

import (
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

func Run(R, D int) game.Game {
	var actions = [4]game.Move{game.Up, game.Down, game.Left, game.Right}
	g := game.Game{}

	for !g.Done {
		scores := map[game.Move]int64{}

		for _, init_action := range actions {
			cg := g.Copy()

			cg.Move(init_action)

			if D == 0 {
				for !cg.Done {
					cg.Move()
				}
			} else {
				for m := 0; m < *D; m++ {
					cg.Move(actions[rnd.Intn(3)])

					if cg.Done {
						break
					}
				}

			}

			scores[action] += cg.Score
		}

		action := getActionFromScores(scores)

		g.Move(action)
	}

	return game
}

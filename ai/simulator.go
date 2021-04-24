package ai

import (
	"github.com/BattlesnakeOfficial/rules"
	"github.com/bmxguy100/battlesnakes_alphabeta/lib"
)

func createSimulation(request lib.GameRequest) (rules.Ruleset, *rules.BoardState) {
	// TODO: Non-solo games
	ruleset := rules.SoloRuleset{
		StandardRuleset: rules.StandardRuleset{
			// Don't depend on food spawning to survive
			// TODO: don't hit own tail from eating food (should the always/sometimes grow?)
			FoodSpawnChance: 0,
			MinimumFood:     0,
		},
	}

	state := rules.BoardState{
		Width:  request.Board.Width,
		Height: request.Board.Height,
	}

	for _, apiFood := range request.Board.Food {
		state.Food = append(state.Food, rules.Point{
			X: apiFood.X,
			Y: apiFood.Y,
		})
	}

	for _, apiSnake := range request.Board.Snakes {
		snake := rules.Snake{
			ID:     apiSnake.ID,
			Health: apiSnake.Health,
		}

		for _, apiBody := range apiSnake.Body {
			snake.Body = append(snake.Body, rules.Point{
				X: apiBody.X,
				Y: apiBody.Y,
			})
		}

		state.Snakes = append(state.Snakes, snake)
	}

	return &ruleset, &state
}

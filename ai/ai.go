package ai

import (
	"github.com/BattlesnakeOfficial/rules"
	api "github.com/BattlesnakeOfficial/rules/cli/commands"
	log "github.com/sirupsen/logrus"
)

const (
	EMPTY = iota
	FOOD
	ENEMY
	YOU
)

var moves = [...]string{"up", "right", "down", "left"}

func SelectMove(request api.ResponsePayload) string {
	ruleset, state := createSimulation(request)

	bestMove := moves[0]
	bestMoveScore := evaluate(iterate(ruleset, state, moves[0]))

	for _, move := range moves[1:] {
		score := evaluate(iterate(ruleset, state, move))
		if score > bestMoveScore {
			bestMove = move
		}
	}

	return bestMove
}

func iterate(ruleset rules.Ruleset, state *rules.BoardState, move string) *rules.BoardState {
	newState, err := ruleset.CreateNextBoardState(state, []rules.SnakeMove{
		{
			ID:   state.Snakes[0].ID,
			Move: move,
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	return newState
}

func evaluate(state *rules.BoardState) (score float32) {
	snake := state.Snakes[0]

	score = 0.0

	if snake.EliminatedCause != "" {
		score = -1e10
		return
	}

	score += float32(snake.Health)

	return
}

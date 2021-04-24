package ai

import (
	"time"

	"github.com/BattlesnakeOfficial/rules"
	"github.com/bmxguy100/battlesnakes_alphabeta/lib"
	log "github.com/sirupsen/logrus"
)

const (
	EMPTY = iota
	FOOD
	ENEMY
	YOU
)

const (
	SCORE_DIE     = -100_000.0
	SCORE_SURVIVE = 10.0
	SCORE_HEALTH  = 1.0
)

var moves = [...]string{"up", "right", "down", "left"}

func SelectMove(request lib.GameRequest) (bestMove string) {
	ruleset, state := createSimulation(request)

	start := time.Now()
	bestMove, bestMoveScore := selectMove(ruleset, state, 10)
	runtime := time.Since(start)

	log.WithField("move", bestMove).
		WithField("score", bestMoveScore).
		WithField("runtime", runtime).
		Info("Chose a move")

	return
}

func selectMove(ruleset rules.Ruleset, state *rules.BoardState, depth int) (bestMove string, bestMoveScore float32) {
	bestMoveScore = 1e-100

	for _, move := range moves {
		var score float32
		newState := iterate(ruleset, state, move)

		score = evaluate(newState)

		if score != SCORE_DIE && depth > 1 {
			_, score = selectMove(ruleset, newState, depth-1)
		}

		if score != SCORE_DIE {
			score += SCORE_SURVIVE
		}

		if score > bestMoveScore {
			bestMove = move
			bestMoveScore = score
		}
	}

	return
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
		return SCORE_DIE
	}

	score += float32(snake.Health) * SCORE_HEALTH

	return
}

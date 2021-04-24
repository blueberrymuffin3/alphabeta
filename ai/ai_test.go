package ai

import (
	"testing"

	battlesnake "github.com/BattlesnakeOfficial/rules/cli/commands"
	log "github.com/sirupsen/logrus"
)

var you = battlesnake.SnakeResponse{
	Id:      "gs_cSGkVSbqhcdVjBVXXgh9kWY3",
	Name:    "blueberry-hackvh",
	Health:  99,
	Shout:   "",
	Squad:   "",
	Latency: 58,
	Head:    battlesnake.Coord{X: 6, Y: 1},
	Length:  23,
	Body: []battlesnake.Coord{
		{X: 6, Y: 1},
		{X: 7, Y: 1},
		{X: 8, Y: 1},
		{X: 8, Y: 2},
		{X: 8, Y: 3},
		{X: 8, Y: 4},
		{X: 8, Y: 5},
		{X: 8, Y: 6},
		{X: 8, Y: 7},
		{X: 8, Y: 8},
		{X: 7, Y: 8},
		{X: 6, Y: 8},
		{X: 6, Y: 7},
		{X: 6, Y: 6},
		{X: 5, Y: 6},
		{X: 5, Y: 5},
		{X: 5, Y: 4},
		{X: 5, Y: 3},
		{X: 5, Y: 2},
		{X: 4, Y: 2},
		{X: 3, Y: 2},
		{X: 2, Y: 2},
		{X: 2, Y: 3},
	},
}

var request = battlesnake.ResponsePayload{
	Turn: 317,
	Game: battlesnake.GameResponse{
		Id:      "sample",
		Timeout: 500,
	},
	Board: battlesnake.BoardResponse{
		Width:  11,
		Height: 11,
		Snakes: []battlesnake.SnakeResponse{you},
		Food: []battlesnake.Coord{
			{X: 10, Y: 2},
			{X: 10, Y: 7},
			{X: 0, Y: 3},
			{X: 9, Y: 2},
			{X: 5, Y: 1},
			{X: 8, Y: 10},
			{X: 7, Y: 0},
			{X: 10, Y: 5},
			{X: 9, Y: 3},
			{X: 10, Y: 4},
			{X: 0, Y: 2},
			{X: 9, Y: 7},
			{X: 3, Y: 1},
			{X: 4, Y: 1},
			{X: 9, Y: 8},
			{X: 2, Y: 0},
			{X: 7, Y: 2},
			{X: 2, Y: 10},
			{X: 5, Y: 10},
			{X: 9, Y: 9},
			{X: 3, Y: 0},
			{X: 9, Y: 4},
			{X: 6, Y: 3},
			{X: 10, Y: 1},
			{X: 7, Y: 7},
			{X: 10, Y: 3},
			{X: 7, Y: 6},
			{X: 0, Y: 1},
			{X: 3, Y: 3},
			{X: 10, Y: 0},
			{X: 6, Y: 10},
			{X: 0, Y: 5},
			{X: 10, Y: 9},
			{X: 0, Y: 7},
		},
		Hazards: []battlesnake.Coord{},
	},
	You: you,
}

func BenchmarkSelectMove(b *testing.B) {
	SelectMove(request)

	log.SetLevel(log.FatalLevel)
	for i := 0; i < b.N-1; i++ {
		SelectMove(request)
	}
	log.SetLevel(log.InfoLevel)
}

func BenchmarkEvaluate(b *testing.B) {
	_, state := createSimulation(request)

	log.SetLevel(log.FatalLevel)
	for i := 0; i < b.N; i++ {
		evaluate(state)
	}
	log.SetLevel(log.InfoLevel)
}

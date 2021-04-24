package ai

import (
	"testing"

	"github.com/bmxguy100/battlesnakes_alphabeta/lib"
	log "github.com/sirupsen/logrus"
)

var you = lib.Battlesnake{
	ID:     "gs_cSGkVSbqhcdVjBVXXgh9kWY3",
	Name:   "blueberry-hackvh",
	Health: 99,
	Shout:  "",
	Head:   lib.Coord{X: 6, Y: 1},
	Length: 23,
	Body: []lib.Coord{
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

var request = lib.GameRequest{
	Turn: 317,
	Game: lib.Game{
		ID:      "sample",
		Timeout: 500,
	},
	Board: lib.Board{
		Width:  11,
		Height: 11,
		Snakes: []lib.Battlesnake{you},
		Food: []lib.Coord{
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

package alphabeta

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bmxguy100/battlesnakes_alphabeta/ai"
	"github.com/bmxguy100/battlesnakes_alphabeta/lib"
	log "github.com/sirupsen/logrus"
)

// handleIndex is called when your Battlesnake is created and refreshed
// by play.battlesnake.com. BattlesnakeInfoResponse contains information about
// your Battlesnake, including what it should look like on the game board.
func handleIndex(w http.ResponseWriter, r *http.Request) {
	response := lib.BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "bmxguy100",
		Color:      "#f2541b",
		Head:       "default",
		Tail:       "default",
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.WithError(err).Error("Error encoding json for index")
		return
	}
}

// handleStart is called at the start of each game your Battlesnake is playing.
// The GameRequest object contains information about the game that's about to start.
// TODO: Use this function to decide how your Battlesnake is going to look on the board.
func handleStart(w http.ResponseWriter, r *http.Request) {
	request := lib.GameRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.WithError(err).Error("Error decoding json for start")
		return
	}

	// Nothing to respond with here
	log.Info("START\n")
}

// handleMove is called for each turn of each game.
// Valid responses are "up", "down", "left", or "right".
// TODO: Use the information in the GameRequest object to determine your next move.
func handleMove(w http.ResponseWriter, r *http.Request) {
	request := lib.GameRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.WithError(err).Warn("Error decoding json for move")
		// return
	}

	move := ai.SelectMove(request)

	response := lib.MoveResponse{
		Move: move,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.WithError(err).Error("Error encoding json for move")
		return
	}
}

// handleEnd is called when a game your Battlesnake was playing has ended.
// It's purely for informational purposes, no response required.
func handleEnd(w http.ResponseWriter, r *http.Request) {
	request := lib.GameRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.WithError(err).Error("Error decoding json for end")
		return
	}

	log.Info("END\n")
}

func handleHealthZ(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

func StartServer(host string, port int64) {
	address := fmt.Sprintf("%s:%d", host, port)

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/start", handleStart)
	http.HandleFunc("/move", handleMove)
	http.HandleFunc("/end", handleEnd)
	http.HandleFunc("/healthz", handleHealthZ)

	log.WithField("port", port).Info("Starting Battlesnake Server")
	log.Fatal(http.ListenAndServe(address, nil))
}

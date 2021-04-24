package alphabeta

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func TestSolo() {
	port := rand.Intn(44150) + 5001
	url := fmt.Sprintf("http://localhost:%d", port)

	go StartServer("localhost", int64(port))

	cmd := exec.Command("battlesnake", "play", "-W", "11", "-H", "11", "--name", "alphabeta", "--url", url, "-g", "solo")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}

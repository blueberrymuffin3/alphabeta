package alphabeta

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"testing"
)

func TestSolo(t *testing.T) {
	port := rand.Intn(44150) + 5001
	url := fmt.Sprintf("http://localhost:%d", port)

	go StartServer("localhost", int64(port))

	cmd := exec.Command("battlesnake", "play", "-W", "11", "-H", "11", "--name", "alphabeta", "--url", url, "-g", "solo")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		t.Error(err)
	}
}

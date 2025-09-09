package server

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	fmt.Fprint(w, getPlayerScore(player))
}

func getPlayerScore(player string) int {
	var score int

	if player == "Floyd" {
		score = 10
	}
	if player == "Pepper" {
		score = 10
	}
	return score
}

package server_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	server "github.com/Youssifahmed12/lgwt-project/httpserver"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}
func TestGETPlayers(t *testing.T) {
	s := &server.PlayerServer{Store: &StubPlayerStore{scores: map[string]int{
		"Pepper": 20,
		"Floyd":  10,
	}}}
	t.Run("returns Pepper's score", func(t *testing.T) {
		req := getNewRequest("Pepper")
		res := httptest.NewRecorder()

		s.ServeHTTP(res, req)

		assertResponseBody(t, res.Body.String(), "20")
		assertStatus(t, res.Code, http.StatusOK)

	})
	t.Run("returns Floyd's score", func(t *testing.T) {
		req := getNewRequest("Floyd")
		res := httptest.NewRecorder()

		s.ServeHTTP(res, req)

		assertResponseBody(t, res.Body.String(), "10")
		assertStatus(t, res.Code, http.StatusOK)

	})

	t.Run("Return 404 on missing players", func(t *testing.T) {
		req := getNewRequest("Youssif")
		res := httptest.NewRecorder()

		s.ServeHTTP(res, req)
		assertStatus(t, res.Code, http.StatusNotFound)
	})
}

func TestStoreScores(t *testing.T) {
	store := &StubPlayerStore{map[string]int{}, nil}
	s := &server.PlayerServer{Store: store}

	player := "Youssif"
	req := newPostWinRequest(player)
	res := httptest.NewRecorder()

	s.ServeHTTP(res, req)

	assertStatus(t, res.Code, http.StatusAccepted)

	if len(store.winCalls) != 1 {
		t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
	}

	if store.winCalls[0] != player {
		t.Errorf("got %q , want %q", store.winCalls[0], player)
	}

}

func newPostWinRequest(player string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
	return req
}

func getNewRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

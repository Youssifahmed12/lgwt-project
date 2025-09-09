package server_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	server "github.com/Youssifahmed12/lgwt-project/httpserver"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}
func TestGETPlayers(t *testing.T) {
	s := &server.PlayerServer{Store: StubPlayerStore{scores: map[string]int{
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

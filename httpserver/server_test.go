package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	server "github.com/Youssifahmed12/lgwt-project/httpserver"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		if err != nil {
			t.Fatal(err)
		}

		res := httptest.NewRecorder()

		server.PlayerServer(res, req)

		got := res.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("returns Floyd's score", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
		if err != nil {
			t.Fatal(err)
		}

		res := httptest.NewRecorder()

		server.PlayerServer(res, req)

		got := res.Body.String()
		want := "10"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

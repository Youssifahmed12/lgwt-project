package main

import (
	"net/http"

	server "github.com/Youssifahmed12/lgwt-project/httpserver"
)

type InMemoryStorage struct {
}

func (I InMemoryStorage) GetPlayerScore(name string) int {
	return 123
}

func (I InMemoryStorage) RecordWin(name string) {

}

func main() {
	s := &server.PlayerServer{Store: InMemoryStorage{}}
	http.ListenAndServe(":5000", s)
}

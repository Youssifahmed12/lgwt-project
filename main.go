package main

import (
	"net/http"

	server "github.com/Youssifahmed12/lgwt-project/httpserver"
)

// in_memory_player_store.go
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func main() {
	s := &server.PlayerServer{Store: NewInMemoryPlayerStore()}
	http.ListenAndServe(":5000", s)
}

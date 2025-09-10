package main

import (
	"net/http"

	server "github.com/Youssifahmed12/lgwt-project/httpserver"
)

func main() {
	s := server.NewPlayerServer(&InMemoryPlayerStore{})
	http.ListenAndServe(":5000", s)
}

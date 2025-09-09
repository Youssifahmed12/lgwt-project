package main

import (
	"net/http"

	server "github.com/Youssifahmed12/lgwt-project/httpserver"
)

func main() {
	handler := http.HandlerFunc(server.PlayerServer)
	http.ListenAndServe(":8080", handler)
}

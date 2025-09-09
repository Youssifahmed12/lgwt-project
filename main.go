package main

import (
	"net/http"

	server "github.com/Youssifahmed12/lgwt-project/httpserver"
)

func main() {
	s := &server.PlayerServer{}
	http.ListenAndServe(":8080", s)
}

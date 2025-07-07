package main

import (
	"github.com/brandoyts/go-clean/infrastructure/rest"
)

func main() {
	srv := rest.NewServer("8000")
	srv.Start()
}

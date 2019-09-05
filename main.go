package main

import (
	"github.com/bobcob7/devplace/internal/server"
	"github.com/bobcob7/devplace/internal/server/stub"
)

func main() {
	handlers := &stub.Stub{}
	server.Start("0.0.0.0:8080", handlers)
}

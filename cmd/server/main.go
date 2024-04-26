package main

import (
	"github.com/mrumyantsev/enigma-app/internal/app/server"
	"github.com/mrumyantsev/logx/log"
)

func main() {
	app, err := server.New()
	if err != nil {
		log.Fatal("failed to initialize application", err)
	}

	if err = app.Run(); err != nil {
		log.Fatal("failed to run application", err)
	}
}

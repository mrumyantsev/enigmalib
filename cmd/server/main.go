package main

import "github.com/mrumyantsev/cipher-machines-app/internal/app/server"

func main() {
	app := server.New()

	app.Run()
}

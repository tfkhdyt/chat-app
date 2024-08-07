package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/tfkhdyt/chat-app/server/lib"
)

func main() {
	lib.NewFx().Run()
}

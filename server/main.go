package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/tfkhdyt/chat-app/server/lib/fx"
)

func main() {
	fx.NewFx().Run()
}

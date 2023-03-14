package main

import (
	"log"

	"github.com/conradgg/product-management/bot/discord"
	"github.com/conradgg/product-management/webhooks"
)

func main() {
	log.Println("Server started")
	LoadConfig()
	go discord.DicordBot()
	webhooks.Receiver()
}

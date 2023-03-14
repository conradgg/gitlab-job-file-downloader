package main

import (
	"log"

	"github.com/conradgg/product-management/webhooks"
)

func main() {
	log.Println("Server started")
	LoadConfig()
	webhooks.Receiver()
}

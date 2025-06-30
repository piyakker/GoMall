package main

import (
	"log"

	config "github.com/piyakker/GoMall/configs"
	"github.com/piyakker/GoMall/router"
)

func main() {

	cfg := config.LoadConfig()

	r := router.SetupRouter(cfg)
	log.Println("API server running at http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

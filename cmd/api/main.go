package main

import (
	"log"

	"github.com/nikhilnarayanan623/go-stripe-payment-gateway/pkg/config"
	"github.com/nikhilnarayanan623/go-stripe-payment-gateway/pkg/di"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("faild to load env, check the .env created and added the values on it according to config")
	}

	server, err := di.InitializeApi(cfg)

	if err != nil {
		log.Fatal("initialize api and server")
	}

	server.Start()
}

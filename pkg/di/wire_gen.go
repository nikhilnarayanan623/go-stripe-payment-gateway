// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/nikhilnarayanan623/go-stripe-payment-gateway/pkg/api"
	"github.com/nikhilnarayanan623/go-stripe-payment-gateway/pkg/api/handler"
	"github.com/nikhilnarayanan623/go-stripe-payment-gateway/pkg/config"
)

// Injectors from wire.go:

func InitializeApi(cfg config.Config) (*http.ServerHTTP, error) {
	stripeHandler := handler.NewStripeHandler(cfg)
	serverHTTP := http.NewServerHTTP(stripeHandler)
	return serverHTTP, nil
}

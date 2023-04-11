//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	http "github.com/nikhilnarayanan623/go-stripe-payment-gateway/pkg/api"
	"github.com/nikhilnarayanan623/go-stripe-payment-gateway/pkg/api/handler"
	"github.com/nikhilnarayanan623/go-stripe-payment-gateway/pkg/config"
)

func InitializeApi(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(
		handler.NewStripeHandler,

		http.NewServerHTTP,
	)

	return &http.ServerHTTP{}, nil
}

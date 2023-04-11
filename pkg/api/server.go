package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/nikhilnarayanan623/ecommerce-gin-clean-arch/cmd/api/docs"
	"github.com/nikhilnarayanan623/go-stripe-payment-gateway/pkg/api/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(stripeHandler *handler.StripeHandler) *ServerHTTP {

	engine := gin.New()

	engine.LoadHTMLGlob("views/*.html")

	engine.Use(gin.Logger())

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//setup end points
	stripe := engine.Group("/stripe")

	stripe.GET("/", stripeHandler.GetStripPaymentPage)
	stripe.POST("/checkout", stripeHandler.StripeCheckout)
	stripe.POST("/verify", stripeHandler.StripeVerify)

	// no route
	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"StatusCode": 404,
			"msg":        "invalid url",
		})
	})

	return &ServerHTTP{
		engine: engine,
	}
}

func (c *ServerHTTP) Start() {
	c.engine.Run(":8000")
}

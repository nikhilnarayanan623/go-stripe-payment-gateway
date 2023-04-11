package handler

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/go-stripe-payment-gateway/pkg/config"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

type StripeHandler struct {
	cfg config.Config
}

// setup stripe handler with env config
func NewStripeHandler(cfg config.Config) *StripeHandler {
	return &StripeHandler{
		cfg: cfg,
	}
}

// for render the stripe payment page
func (c *StripeHandler) GetStripPaymentPage(ctx *gin.Context) {

	ctx.HTML(200, "index.html", nil)
}

// for checkout the stripe to create payment
func (c *StripeHandler) StripeCheckout(ctx *gin.Context) {

	//!! now here payment amount take from form acutally we need to take from databas order amount
	amountToPay, err := strconv.Atoi(ctx.Request.PostFormValue("amount"))

	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(400, gin.H{
			"message": "invalid amount",
			"error":   err.Error(),
		})
		return
	}

	// set up the stip secret key
	stripe.Key = config.GetCofig().StripSecretKey

	// create a payment param
	params := &stripe.PaymentIntentParams{

		Amount: stripe.Int64(int64(amountToPay)),
		//ReceiptEmail: stripe.String(recieptEmail),

		Currency: stripe.String(string(stripe.CurrencyINR)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	// if an error shows on here change the import packge of paymentIntent to ("github.com/stripe/stripe-go/v72/paymentintent")
	// and also check both stripe import package name inclued v72
	// creata new payment intent with this param
	paymentIntent, err := paymentintent.New(params)

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": fmt.Errorf("faild to create strip payment for amount %v", amountToPay),
		})
		return
	}

	clientSecret := paymentIntent.ClientSecret
	publishableKey := config.GetCofig().StripPublishKey

	ctx.JSON(200, gin.H{
		"amount_to_pay":   amountToPay,
		"publishable_key": publishableKey,
		"client_secret":   clientSecret,
	})
}

// for verify the payment on backend
func (c *StripeHandler) StripeVerify(ctx *gin.Context) {

	paymentID := ctx.Request.PostFormValue("payment_id")
	fmt.Println(paymentID, "pid")

	// set the stripeKey
	stripe.Key = config.GetCofig().StripSecretKey
	// get payment by payment_id
	paymentIntent, err := paymentintent.Get(paymentID, nil)
	// fmt.Println("paymentIntent", paymentIntent)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": fmt.Errorf("faild to get stripe paymentIntent of payment_id %v", paymentID),
		})
		return
	}

	// verify the payment intent
	if paymentIntent.Status != stripe.PaymentIntentStatusSucceeded && paymentIntent.Status != stripe.PaymentIntentStatusRequiresCapture {

		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "payment not not completed",
			})
			return
		}
	}

	ctx.JSON(200, gin.H{
		"error":   nil,
		"message": "successfully payment vefied",
	})
}

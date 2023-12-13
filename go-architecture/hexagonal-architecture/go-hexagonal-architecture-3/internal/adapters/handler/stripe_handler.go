package handler

import (
	"Hexagonal-Architecture/internal/adapters/repository"
	"Hexagonal-Architecture/internal/core/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
)

type PaymentHandler struct {
	svc services.PaymentService
}

func NewPaymentHandler(paymentService services.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		svc: paymentService,
	}
}

func (h *PaymentHandler) CreateCheckoutSession(ctx *gin.Context) {
	apiCfg, err := repository.LoadAPIConfig()
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}
	stripe.Key = apiCfg.StripeKey

	domain := "http://localhost:4242"
	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				// Provide the exact Price ID (for example, pr_1234) of the product you want to sell
				Price:    stripe.String("price_1N5VNbKb78q3bJ6obePPkame"),
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(domain + "?success=true"),
		CancelURL:  stripe.String(domain + "?canceled=true"),
	}

	s, err := session.New(params)

	if err != nil {
		log.Printf("session.New: %v", err)
	}
	orderID := generateOrderID()

	// Add the order ID to the checkout session metadata
	s.Metadata = map[string]string{
		"order_id": orderID,
	}

	ctx.Redirect(http.StatusSeeOther, s.URL)
}

func generateOrderID() string {
	// Generate a unique order ID, for example:
	return uuid.New().String()
}

func (h *PaymentHandler) HandleSuccess(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "<h1>Payment Successful!</h1>", nil)
}

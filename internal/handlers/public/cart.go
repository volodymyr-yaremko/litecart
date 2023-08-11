package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"

	"github.com/shurco/litecart/internal/models"
	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/pkg/security"
	"github.com/shurco/litecart/pkg/webutil"
)

// Checkout is ...
// [post] /cart/checkout
func Checkout(c *fiber.Ctx) error {
	items := &[]models.CheckoutLineItem{}
	if err := c.BodyParser(items); err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	db := queries.DB()
	settingStripe, err := db.SettingStripe()
	if err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	stripe.Key = settingStripe.SecretKey

	lineItems := []*stripe.CheckoutSessionLineItemParams{}
	for _, item := range *items {
		lineItems = append(lineItems, &stripe.CheckoutSessionLineItemParams{
			Price:    stripe.String(item.Price),
			Quantity: stripe.Int64(int64(item.Quantity)),
		})
	}

	cartID := security.RandomString()
	params := &stripe.CheckoutSessionParams{
		LineItems: lineItems,
		//AutomaticTax: &stripe.CheckoutSessionAutomaticTaxParams{
		//	Enabled: stripe.Bool(true),
		//},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(settingStripe.Domain + "/cart/success/" + cartID + "/{CHECKOUT_SESSION_ID}"),
		CancelURL:  stripe.String(settingStripe.Domain + "/cart/cancel/" + cartID),
	}

	stripeSession, err := session.New(params)
	if err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	db.AddCart(&models.Checkout{
		ID:            cartID,
		Cart:          *items,
		AmountTotal:   stripeSession.AmountTotal,
		Currency:      string(stripeSession.Currency),
		PaymentStatus: string(stripeSession.PaymentStatus),
	})

	return webutil.Response(c, fiber.StatusOK, "Checkout url", stripeSession.URL)
}

// CheckoutSuccess is ...
// [get] /cart/success/:cart_id/:session_id
func CheckoutSuccess(c *fiber.Ctx) error {
	db := queries.DB()

	cartID := c.Params("cart_id")
	sessionID := c.Params("session_id")

	sessionStripe, err := session.Get(sessionID, nil)
	if err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	err = db.UpdateCart(&models.Checkout{
		ID:            cartID,
		Email:         sessionStripe.CustomerDetails.Email,
		Name:          sessionStripe.CustomerDetails.Name,
		PaymentID:     sessionStripe.PaymentIntent.ID,
		PaymentStatus: string(sessionStripe.PaymentStatus),
	})
	if err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	return c.Render("site/success", nil, "site/layouts/main")
}

// CheckoutCancel is ...
// [get] /cart/cancel/:cart_id
func CheckoutCancel(c *fiber.Ctx) error {
	cartID := c.Params("cart_id")
	db := queries.DB()
	err := db.UpdateCart(&models.Checkout{
		ID:            cartID,
		PaymentStatus: "cancel",
	})
	if err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	return c.Render("site/cancel", nil, "site/layouts/main")
}

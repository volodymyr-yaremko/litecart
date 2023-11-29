package litepay

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type stripe struct {
	Cfg
	apiToken   string
	successURL string
	cancelURL  string
}

func (c Cfg) Stripe(apiToken string) LitePay {
	c.paymentSystem = STRIPE
	c.api = "https://api.stripe.com"
	c.currency = []string{"EUR", "USD", "GBP", "AUD", "CAD", "JPY", "CNY", "SEK"}
	return &stripe{
		Cfg:        c,
		apiToken:   apiToken,
		successURL: c.successURL,
		cancelURL:  c.cancelURL,
	}
}

func (c *stripe) Pay(cart Cart) (*Payment, error) {
	currency := strings.ToUpper(cart.Currency)
	if !findInSlice(c.currency, strings.ToUpper(currency)) {
		return nil, errors.New("this currency is not supported")
	}

	params := url.Values{}
	for i, s := range cart.Items {
		iString := strconv.Itoa(i)
		params.Add("line_items["+iString+"][price_data][unit_amount]", strconv.Itoa(s.PriceData.UnitAmount))
		params.Add("line_items["+iString+"][price_data][currency]", currency)
		params.Add("line_items["+iString+"][price_data][product_data][name]", s.PriceData.Product.Name)
		for ii, img := range s.PriceData.Product.Images {
			params.Add("line_items["+iString+"][price_data][product_data][images]["+strconv.Itoa(ii)+"]", img)
		}
		params.Add("line_items["+iString+"][quantity]", strconv.Itoa(s.Quantity))
	}
	params.Add("success_url", fmt.Sprintf("%s/?payment_system=%s&cart_id=%s&session={CHECKOUT_SESSION_ID}", c.successURL, c.paymentSystem, cart.ID))
	params.Add("cancel_url", fmt.Sprintf("%s/?payment_system=%s&cart_id=%s", c.cancelURL, c.paymentSystem, cart.ID))
	params.Add("mode", `payment`)
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest(
		http.MethodPost,
		c.api+"/v1/checkout/sessions",
		body,
	)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.apiToken, "")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := parseBody(resp.Body)
	if err != nil {
		return nil, err
	}

	checkout := &Payment{
		AmountTotal:   int(data["amount_total"].(float64)),
		Currency:      strings.ToUpper(data["currency"].(string)),
		Status:        StatusPayment(STRIPE, data["payment_status"].(string)),
		URL:           data["url"].(string),
		PaymentSystem: c.paymentSystem,
	}

	return checkout, nil
}

func (c *stripe) Checkout(payment *Payment, session string) (*Payment, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		c.api+"/v1/checkout/sessions/"+session,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.apiToken, "")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("The server returned an error.")
	}

	data, err := parseBody(resp.Body)
	if err != nil {
		return nil, err
	}

	payment.MerchantID = data["payment_intent"].(string)
	payment.AmountTotal = int(data["amount_total"].(float64))
	payment.Currency = strings.ToUpper(data["currency"].(string))
	payment.Status = StatusPayment(STRIPE, data["payment_status"].(string))

	return payment, nil
}

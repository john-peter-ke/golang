package factory

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrototype(t *testing.T) {
	stripe := NewPaymentClient(&StripeFactory{})
	paypal := NewPaymentClient(&PayPalFactory{})

	assert.Equal(t, true, strings.Contains(stripe.Process(float64(100.00)), "Stripe"))
	assert.Equal(t, true, strings.Contains(paypal.Process(float64(100.00)), "PayPal"))
}

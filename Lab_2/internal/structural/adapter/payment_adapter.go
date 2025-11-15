package adapter

import (
    "fmt"
    "math"
)

type PaymentProvider interface {
    Pay(amount float64) error
}

type LegacyPaymentGateway struct{ MerchantID string }

func (g *LegacyPaymentGateway) MakePayment(cents int64) error {
    fmt.Printf("[LegacyPayment] Charging %d cents via merchant %s\n", cents, g.MerchantID)
    return nil
}

type LegacyPaymentAdapter struct{ Gateway *LegacyPaymentGateway }

func (a *LegacyPaymentAdapter) Pay(amount float64) error {
    cents := int64(math.Round(amount * 100))
    return a.Gateway.MakePayment(cents)
}

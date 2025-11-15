package main

import (
    "fmt"
    "example.com/bloomify-go/internal/domain"
    "example.com/bloomify-go/internal/structural/adapter"
    "example.com/bloomify-go/internal/structural/decorator"
    "example.com/bloomify-go/internal/structural/facade"
)

func main() {
    baseBouquet := &domain.BasicBouquet{Name: "Romantic Roses", BasePrice: 350}
    decorated := &decorator.WithVase{Wrapped: &decorator.WithCard{Wrapped: &decorator.WithRibbon{Wrapped: baseBouquet}}}

    gateway := &adapter.LegacyPaymentGateway{MerchantID: "BLOOMIFY-123"}
    provider := &adapter.LegacyPaymentAdapter{Gateway: gateway}

    orderService := facade.NewOrderService(provider)
    orderService.PlaceOrder(decorated, "Popescu Sabina")

    fmt.Println("Done.")
}

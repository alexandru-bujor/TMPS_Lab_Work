package factory

import "fmt"

type PaymentMethod interface {
    Pay(amount float64)
}

type CardPayment struct{}
func (c *CardPayment) Pay(amount float64) {
    fmt.Println("Paid with card:", amount)
}

type CashPayment struct{}
func (c *CashPayment) Pay(amount float64) {
    fmt.Println("Paid with cash:", amount)
}

func GetPaymentMethod(method string) PaymentMethod {
    if method == "card" {
        return &CardPayment{}
    }
    return &CashPayment{}
}

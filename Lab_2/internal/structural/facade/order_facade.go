package facade

import (
	"TMPS_Lab_Work/Lab_2/internal/domain"
	"TMPS_Lab_Work/Lab_2/internal/structural/adapter"
	"fmt"
)

type OrderService struct{ provider adapter.PaymentProvider }

func NewOrderService(p adapter.PaymentProvider) *OrderService { return &OrderService{provider: p} }

func (s *OrderService) PlaceOrder(b domain.Bouquet, customer string) error {
	fmt.Println("=== ORDER ===")
	fmt.Println("Customer:", customer)
	fmt.Println("Bouquet:", b.Description())
	fmt.Println("Price:", b.Price())
	s.provider.Pay(b.Price())
	fmt.Println("Order completed.")
	return nil
}

package main

import (
    "fmt"
    "example.com/creational/internal/creational/singleton"
    "example.com/creational/internal/creational/builder"
    "example.com/creational/internal/creational/factory"
)

func main() {
    // Singleton
    config := singleton.GetConfig()
    fmt.Println("Config value:", config.Value)

    // Builder
    bouquet := builder.NewBouquetBuilder().SetName("Spring Mix").AddFlower("Rose").AddFlower("Tulip").Build()
    fmt.Println("Bouquet:", bouquet.Name, bouquet.Flowers)

    // Factory
    payment := factory.GetPaymentMethod("card")
    payment.Pay(120)
}

package main

import (
	"TMPS_Lab_Work/Lab_3/internal/behavioral/command"
	"TMPS_Lab_Work/Lab_3/internal/behavioral/observer"
	"TMPS_Lab_Work/Lab_3/internal/behavioral/strategy"
	"TMPS_Lab_Work/Lab_3/internal/domain"
	"fmt"
)

func main() {
	// Observer
	order := &observer.Order{}
	order.Attach(&observer.EmailNotifier{})
	order.SetStatus("Packed")

	// Strategy
	ctx := &strategy.Context{}
	ctx.SetStrategy(&strategy.Drone{})
	ctx.Execute(1)

	// Command
	b := &domain.Bouquet{Name: "Roses"}
	mgr := &command.Manager{}
	add := &command.AddFlower{Bouquet: b, Flower: "Rose"}
	mgr.Run(add)
	fmt.Println("After add:", b.Flowers)
	mgr.Undo()
	fmt.Println("After undo:", b.Flowers)
}

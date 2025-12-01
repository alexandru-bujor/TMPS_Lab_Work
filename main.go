package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	lab0solid "TMPS_Lab_Work/Lab_0"
)

func main() {
	fmt.Println("=== Labs Runner ===")
	fmt.Println("Available labs:")
	fmt.Println("  1) Laboratory Work #1 – SOLID (SRP, OCP, DIP)")
	fmt.Print("Enter lab number to run (e.g., 1): ")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	choice := strings.TrimSpace(text)
	if choice == "" {
		choice = "1" // default
	}

	switch choice {
	case "1":
		fmt.Println("\nRunning Laboratory Work #1 – SOLID (SRP, OCP, DIP)\n")
		lab0solid.Run()
	default:
		fmt.Println("Unknown lab. Available: 1")
	}
}

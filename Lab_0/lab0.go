package Lab_0

import "fmt"

// Run is the public entrypoint for Lab 1.
// It prints explanations for each SOLID principle, then demos them.
func Run() {
	// 1) SRP
	fmt.Println("—— SRP (Single Responsibility Principle) ——")

	report := Report{Title: "Lab 1 - SOLID", Text: "Understanding SRP, OCP, and DIP"}
	report.Display()
	fmt.Println()

	// 2) OCP
	fmt.Println("—— OCP (Open/Closed Principle) ——")

	circle := Circle{Radius: 3}
	square := Square{Side: 4}
	PrintArea(circle)
	PrintArea(square)
	fmt.Println()

	// 3) DIP
	fmt.Println("—— DIP (Dependency Inversion Principle) ——")

	notifyEmail := Notification{Sender: EmailSender{}}
	notifySMS := Notification{Sender: SMSSender{}}
	notifyEmail.Alert()
	notifySMS.Alert()

	fmt.Println("\n✅ Lab 1 finished:  SRP (Report), OCP (Shape), DIP (Notification/MessageSender).")
}

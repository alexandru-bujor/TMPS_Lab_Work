package Lab_0

import "fmt"

type MessageSender interface {
	Send(msg string)
}

type EmailSender struct{}

func (EmailSender) Send(msg string) { fmt.Println("📧 Sending email:", msg) }

type SMSSender struct{}

func (SMSSender) Send(msg string) { fmt.Println("📱 Sending SMS:", msg) }

type Notification struct {
	Sender MessageSender
}

func (n Notification) Alert() {
	n.Sender.Send("System alert triggered!")
}

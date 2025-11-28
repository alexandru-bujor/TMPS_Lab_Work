package observer

type Observer interface { Update(status string) }

type Subject interface {
    Attach(o Observer)
    Notify()
}

type Order struct {
    Status    string
    observers []Observer
}

func (o *Order) Attach(obs Observer) { o.observers = append(o.observers, obs) }
func (o *Order) Notify() { for _, ob := range o.observers { ob.Update(o.Status) } }
func (o *Order) SetStatus(s string) { o.Status = s; o.Notify() }

type EmailNotifier struct{}
func (e *EmailNotifier) Update(status string) {}

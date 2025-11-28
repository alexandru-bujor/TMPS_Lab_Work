
# **Behavioral Design Patterns –(Go Implementation)**
### **Author:** Bujor Alexandru
### **Course:** TMPS – Software Design Techniques
### **Laboratory Work – Behavioral Design Patterns**

---

##  **1. Introduction / Motivation**

This laboratory work continues the development of the **Bloomify** flower shop system by integrating *Behavioral Design Patterns*.  
While previous labs introduced Creational and Structural patterns, the focus here is on improving **communication**, **interaction**, and **runtime behavior** between software components.

Behavioral Design Patterns help the system:

- react dynamically to changes,
- encapsulate operations and algorithms,
- improve flexibility,
- add reversible actions,
- support interchangeable behaviors,
- and reduce coupling between modules.

For Bloomify, the following behavioral needs were identified:

- Orders change status and multiple components should react → **Observer**
- Delivery methods differ depending on user choice → **Strategy**
- Admin operations should allow undo/redo → **Command**

These patterns enrich Bloomify’s functionality and align perfectly with the system’s practical requirements.

---

##  **2. Theory – Behavioral Design Patterns**

Behavioral patterns define **communication mechanisms** between entities. Unlike creational or structural patterns, which focus on object creation or composition, behavioral patterns define *how* objects interact.

Examples include:

- Observer
- Strategy
- Command
- Mediator
- Chain of Responsibility
- Interpreter
- Iterator

This project implements:

### ✔ **Observer Pattern**
Used when multiple dependent objects must react to changes in another object.

### ✔ **Strategy Pattern**
Used when multiple algorithms can be swapped at runtime.

### ✔ **Command Pattern**
Used when operations should be encapsulated as objects with reversible behavior.

---

#  **3. Implemented Behavioral Patterns**

---

## **3.1 Observer Pattern – Order Status Notifications**

### **Motivation**
Bloomify must notify customers or internal systems whenever an order changes status:

- "Created"
- "Packed"
- "Out for Delivery"
- "Delivered"

Instead of hardcoding notification logic inside the `Order` struct, **Observer** separates responsibilities and allows dynamic subscription.

---

###  **Location:**
`internal/behavioral/observer/observer.go`

###  **Code Snippet**

```go
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
func (e *EmailNotifier) Update(status string) {
    fmt.Println("[Email] Order status changed:", status)
}
```

### **Why This Pattern Fits**
- Multiple observers can subscribe independently
- Decouples notification logic from order management
- Extensible (add SMSNotifier, PushNotifier, etc.)

---

## **3.2 Strategy Pattern – Delivery Method Selection**

### **Motivation**
Bloomify supports various delivery modes:

- Courier delivery
- In-store pickup
- Autonomous drone delivery

Each delivery method has *different behavior*, but the client should interact through a single interface.

---

###  **Location:**
`internal/behavioral/strategy/strategy.go`

###  **Code Snippet**

```go
type DeliveryStrategy interface { Deliver(orderID int) }

type Courier struct{}
func (c *Courier) Deliver(id int) { fmt.Println("Courier delivers", id) }

type Drone struct{}
func (d *Drone) Deliver(id int) { fmt.Println("Drone delivers", id) }

type Context struct { Strategy DeliveryStrategy }

func (c *Context) SetStrategy(s DeliveryStrategy) { c.Strategy = s }
func (c *Context) Execute(id int) { c.Strategy.Deliver(id) }
```

### ✔ **Why It Fits**
- Delivery algorithm can change at runtime
- Supports adding new methods without modifying existing code
- Reduces conditional logic

---

## **3.3 Command Pattern – Undo/Redo for Admin Actions**

### **Motivation**
Administrators modify bouquets frequently:

- Add flower
- Remove flower
- Update price

These operations should be reversible.

**Command Pattern** encapsulates each action as an object with `Execute()` and `Undo()`.

---

###  **Location:**
`internal/behavioral/command/command.go`

###  **Code Snippet**

```go
type Command interface {
    Execute()
    Undo()
}

type AddFlower struct {
    Bouquet *domain.Bouquet
    Flower  string
}

func (c *AddFlower) Execute() { c.Bouquet.Flowers = append(c.Bouquet.Flowers, c.Flower) }
func (c *AddFlower) Undo() { c.Bouquet.Flowers = c.Bouquet.Flowers[:len(c.Bouquet.Flowers)-1] }

type Manager struct{ History []Command }

func (m *Manager) Run(cmd Command) {
    cmd.Execute()
    m.History = append(m.History, cmd)
}

func (m *Manager) Undo() {
    if len(m.History) == 0 { return }
    last := m.History[len(m.History)-1]
    last.Undo()
    m.History = m.History[:len(m.History)-1]
}
```

###  **Why This Pattern Fits**
- Allows full history of operations
- Enables Undo/Redo
- Clean separation of user actions from logic

---

#  **4. Demonstration (Client Program)**

###  **Location:**
`cmd/client/main.go`

###  **Code Snippet**

```go
order := &observer.Order{}
order.Attach(&observer.EmailNotifier{})
order.SetStatus("Packed")

ctx := &strategy.Context{}
ctx.SetStrategy(&strategy.Drone{})
ctx.Execute(1)

b := &domain.Bouquet{Name: "Roses"}
mgr := &command.Manager{}
add := &command.AddFlower{Bouquet: b, Flower: "Rose"}

mgr.Run(add)
fmt.Println("After add:", b.Flowers)

mgr.Undo()
fmt.Println("After undo:", b.Flowers)
```

---

#  **5. Results (Sample Output)**

```
[Email] Order status changed: Packed
Drone delivers 1
After add: [Rose]
After undo: []
```

---

#  **6. Project Structure**

```

README.md
cmd/client/main.go
internal/domain/bouquet.go
internal/behavioral/
  observer/
  observer.go
  strategy/
    strategy.go
  command/
    command.go
```

---

#  **7. Conclusions**

This laboratory work successfully demonstrates three core Behavioral Design Patterns integrated into the Bloomify system:

### ✔ Observer – event-driven updates
### ✔ Strategy – interchangeable delivery algorithms
### ✔ Command – reversible admin actions

These improvements provide:

- High flexibility
- Cleaner communication between components
- Extensibility for new features
- Low coupling and maintainable architecture

Together with Creational and Structural patterns, Bloomify now represents a complete, well-architected TMPS project.

---

#  **Lab Requirements Completed**
- Behavioral DP implemented
- Code structured into required directories
- Integrated with previously created architecture
- README documenting theory + implementation
- One client entry point  


# **Structural Design Patterns – (Go Implementation)**
### *Author: Bujor Alexandru*
### *Course: TMPS – Software Design Techniques*
### *Laboratory Work – Structural Design Patterns*

---

## **1. Introduction**

This laboratory work focuses on applying **Structural Design Patterns** to extend the Bloomify flower‑shop system.  
While the previous laboratory explored *Creational* patterns, this assignment emphasizes *how objects are combined*, how subsystems interact, and how complexity can be reduced through well‑designed abstractions.

The goal is to integrate **three structural design patterns** in a clean Go project:

- **Decorator** – add bouquet features at runtime
- **Adapter** – integrate a legacy payment gateway
- **Facade** – provide a simplified unified API (`OrderService`)

Bloomify is an ideal domain because bouquet customization, third‑party service integration, and user‑facing APIs naturally require structural flexibility.
(Bloomify is the template the new team for PBL used as their project front. This is where the ideea comes from.)
---

## **2. Theory Overview – Structural Design Patterns**

Structural patterns define *how classes and objects are composed* to form larger, more flexible architectures.

### Why use structural patterns?

- To reduce coupling between components
- To unify incompatible interfaces
- To extend behavior without modifying core structures
- To simplify complex subsystems
- To promote modularity and clean architecture

### Patterns used in this project:

| Pattern | Purpose | Why Used Here |
|--------|---------|----------------|
| **Decorator** | Attach responsibilities dynamically | Add ribbon, card, vase to bouquets at runtime |
| **Adapter** | Convert interface of one class to another | Integrate a legacy payment API |
| **Facade** | Provide a unified, simple interface | Expose Bloomify operations via `OrderService` |

---

## **3. Domain Description – Bloomify Flower Shop**

Bloomify allows users to:

- Select and customize bouquets
- Add decorative options
- Pay for their order
- Place the final order through a single API

### UML‑like diagram

```
+------------------+
|   Bouquet         |
+------------------+
| Description()     |
| Price()           |
+---------+---------+
          |
    Decorators wrap this
```

```
+--------------------------+
| PaymentProvider          |
+--------------------------+
| Pay(amount float64)      |
+------------+-------------+
             |
    Adapter implements this
```

```
+------------------------+
|     OrderService       |
+------------------------+
| PlaceOrder(...)        |
+------------------------+
   Facade over the whole system
```

---

# **4. Implemented Structural Design Patterns**

---

# **4.1 Decorator Pattern — Bouquet Customization**

### **Motivation**
Bouquets need dynamic customization:

- Ribbon
- Greeting card
- Vase

Creating subclasses like:

- `BouquetWithRibbonAndCard`
- `BouquetWithCardAndVase`

…would explode combinatorially.

**Decorator solves this by using composition**, adding features at runtime.

---

### **Implementation: `bouquet_decorator.go`**

```go
type WithRibbon struct{ Wrapped domain.Bouquet }
func (d *WithRibbon) Description() string { return d.Wrapped.Description() + ", with ribbon" }
func (d *WithRibbon) Price() float64      { return d.Wrapped.Price() + 3.5 }

type WithCard struct{ Wrapped domain.Bouquet }
func (d *WithCard) Description() string { return d.Wrapped.Description() + ", with card" }
func (d *WithCard) Price() float64      { return d.Wrapped.Price() + 5 }

type WithVase struct{ Wrapped domain.Bouquet }
func (d *WithVase) Description() string { return d.Wrapped.Description() + ", with vase" }
func (d *WithVase) Price() float64      { return d.Wrapped.Price() + 25 }
```

### **Usage Example**

```go
bouquet := &domain.BasicBouquet{Name: "Romantic Roses", BasePrice: 350}

custom := &decorator.WithVase{
    Wrapped: &decorator.WithCard{
        Wrapped: &decorator.WithRibbon{Wrapped: bouquet},
    },
}

fmt.Println(custom.Description(), custom.Price())
```

### **Benefits**
- Adds features without modifying the base class
- Flexible runtime combinations
- Clean and maintainable

---

# **4.2 Adapter Pattern — Legacy Payment Gateway Integration**

### **Motivation**
Bloomify integrates a *legacy payment provider* whose API does not match the system’s needs.

We want this clean interface:

```go
type PaymentProvider interface {
    Pay(amount float64) error
}
```

But the legacy API uses:

```go
MakePayment(cents int64)
```

The **Adapter converts between these two worlds.**

---

### **Implementation: `payment_adapter.go`**

```go
type LegacyPaymentGateway struct {
    MerchantID string
}

func (g *LegacyPaymentGateway) MakePayment(totalInCents int64) error {
    fmt.Printf("[LegacyPayment] Charging %d cents via merchant %s\n", totalInCents, g.MerchantID)
    return nil
}

type LegacyPaymentAdapter struct{ Gateway *LegacyPaymentGateway }

func (a *LegacyPaymentAdapter) Pay(amount float64) error {
    cents := int64(math.Round(amount * 100))
    return a.Gateway.MakePayment(cents)
}
```

### **Benefits**
- Allows reuse of old code without modification
- System depends only on high‑level interfaces
- Future payment providers can be swapped easily

---

# **4.3 Facade Pattern — High‑Level Order API**

### **Motivation**
Bloomify involves several steps:

- Configure bouquet
- Apply decorators
- Process payment
- Send notifications
- Display results

Controllers should not deal with these details.  
The **Facade** provides a **single, simple entry point**.

---

### **Implementation: `order_facade.go`**

```go
type OrderService struct{ provider adapter.PaymentProvider }

func NewOrderService(p adapter.PaymentProvider) *OrderService {
    return &OrderService{provider: p}
}

func (s *OrderService) PlaceOrder(b domain.Bouquet, customer string) error {
    fmt.Println("=== ORDER ===")
    fmt.Println("Customer:", customer)
    fmt.Println("Bouquet:", b.Description())
    fmt.Println("Price:", b.Price())

    s.provider.Pay(b.Price())

    fmt.Println("Order completed.")
    return nil
}
```

### **Benefits**
- Hides internal complexity
- Reduces coupling
- Clean separation between *client* and *subsystems*

---

# **5. Client Demonstration**

### File: `main.go`

```go
baseBouquet := &domain.BasicBouquet{Name: "Romantic Roses", BasePrice: 350}

decorated := &decorator.WithVase{
    Wrapped: &decorator.WithCard{
        Wrapped: &decorator.WithRibbon{Wrapped: baseBouquet},
    },
}

gateway := &adapter.LegacyPaymentGateway{MerchantID: "BLOOMIFY-123"}
provider := &adapter.LegacyPaymentAdapter{Gateway: gateway}

orderService := facade.NewOrderService(provider)
orderService.PlaceOrder(decorated, "Bujor Alexandru")
```

### Output Example

```
=== ORDER ===
Customer: Bujor Alexandru
Bouquet: Romantic Roses, with ribbon, with card, with vase
Price: 383.50
[LegacyPayment] Charging 38350 cents via merchant BLOOMIFY-123
Order completed.
```

---

# **6. Project Structure**

```

cmd/client/main.go
internal/domain/bouquet.go
internal/structural/
  decorator/
    bouquet_decorator.go
  adapter/
    payment_adapter.go
  facade/
    order_facade.go
README.md
```

Matches TMPS requirements for modular structure.

---

#  **7. Conclusions**

This laboratory work demonstrates the power of Structural Design Patterns:

### ✔ Decorator
Flexible bouquet customization using composition.

### ✔ Adapter
Smooth integration with incompatible external systems.

### ✔ Facade
A simple, unified interface for complex operations.

### Result


- Clean architecture
- Low coupling
- High flexibility
- Extensibility for future features


---


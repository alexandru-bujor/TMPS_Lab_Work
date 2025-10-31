
# **Creational Design Patterns – Bloomify (Go Implementation)**
### *Author: Bujor Alexandru, FAF-231*
### *Course: TMPS – Software Design Techniques*
### *Laboratory Work – Creational Design Patterns*

---

##  **1. Introduction**

This laboratory work focuses on understanding and applying **Creational Design Patterns** within an object-oriented architectural context. These patterns abstract and improve object creation, ensuring that system components remain flexible, scalable, and easy to maintain.

For this assignment, the chosen domain is **Bloomify**, a simplified digital flower shop system. The domain naturally requires controlled object creation:

- Bouquets vary in structure and composition
- Payment methods differ based on context
- A central configuration must remain consistent throughout the system

Because of this, the domain is well-suited for demonstrating multiple Creational Design Patterns.

This project uses **Go (Golang)** without external frameworks — aligned with the lab constraints.

---

##  **2. Theory Overview – Creational Design Patterns**

Creational Design Patterns deal with controlling *how* objects are created.

### Why control object creation?

- Direct `new` / constructor calls become repetitive
- Complex objects need step-by-step assembly
- Global shared resources must ensure consistency
- Systems must adapt to different creation conditions
- Switching creation strategies should not break existing code

Creational patterns provide structured, safe, and maintainable instantiation strategies.

### Patterns used in this project:

| Pattern | Purpose | Why Used Here |
|--------|---------|----------------|
| **Singleton** | One shared instance globally | A shared configuration for the Bloomify system |
| **Builder** | Build complex objects step-by-step | Bouquets consist of multiple flowers and attributes |
| **Factory Method** | Choose which concrete object to instantiate | Selecting payment types at runtime |

---

##  **3. Domain Description – Bloomify Flower Shop**

### Main Entities:
- **Bouquet** — product containing flowers
- **Configuration** — global shared settings (currency, shop name, etc.)
- **Payment Methods** — card or cash

### UML-like entity sketch:

```
+------------------+        +------------------+
|   Configuration  |        |     Bouquet      |
+------------------+        +------------------+
| Value: string    |        | Name: string     |
| (Singleton)      |        | Flowers: []string|
+------------------+        +------------------+

+------------------+
| PaymentMethod    |
+------------------+
| Pay(amount)      |
+---------+--------+
          |
   -------------------------
   |                       |
+--------+           +-----------+
| Card   |           | Cash      |
+--------+           +-----------+
(Factory Method decides which one)
```

---

# ️ **4. Implemented Creational Patterns**

---

# **4.1 Singleton Pattern — Global Configuration**

### **Motivation**
Bloomify requires a single set of shared settings:  
theme, currency, environment flags, etc.

A Singleton ensures:

- Only **one instance** exists
- System components always refer to the same configuration
- Lazy initialization when needed

### **Code: `internal/creational/singleton/config.go`**

```go
package singleton

type configuration struct {
    Value string
}

var instance *configuration

// GetConfig returns the single configuration instance.
func GetConfig() *configuration {
    if instance == nil {
        instance = &configuration{Value: "Bloomify Default Config"}
    }
    return instance
}
```

### **Usage in client:**

```go
config := singleton.GetConfig()
fmt.Println("Config value:", config.Value)
```

### **Why this works**
- The `instance` variable is private to the package.
- Initialization happens once ("lazy loading").
- Further calls return the same object reference.

---

# **4.2 Builder Pattern — Constructing Complex Bouquets**

### **Motivation**
Bouquets can be:

- Empty
- Contain any number of flowers
- Easily extended with new components later

Hardcoding constructors like:

```go
NewBouquet("Spring Mix", []string{"Rose", "Tulip", "Lily"})
```

makes code unreadable and rigid.

### **Builder solves this**
It allows:

- Step-by-step configuration
- Chainable method calls
- Clear separation of construction and representation

### **Code: `internal/creational/builder/bouquet_builder.go`**

```go
package builder

import "example.com/creational/internal/domain"

type BouquetBuilder struct {
    name    string
    flowers []string
}

func NewBouquetBuilder() *BouquetBuilder {
    return &BouquetBuilder{}
}

func (b *BouquetBuilder) SetName(name string) *BouquetBuilder {
    b.name = name
    return b
}

func (b *BouquetBuilder) AddFlower(f string) *BouquetBuilder {
    b.flowers = append(b.flowers, f)
    return b
}

func (b *BouquetBuilder) Build() domain.Bouquet {
    return domain.Bouquet{Name: b.name, Flowers: b.flowers}
}
```

### **Usage in client:**

```go
bouquet := builder.NewBouquetBuilder().
    SetName("Spring Mix").
    AddFlower("Rose").
    AddFlower("Tulip").
    Build()

fmt.Println(bouquet.Name, bouquet.Flowers)
```

### **Benefits**
- Easily extendable
- No need for complex constructors
- Ideal for assembling many pieces incrementally

---

# **4.3 Factory Method — Dynamic Payment Method Selection**

### **Motivation**
Payment methods differ depending on customer choice:

- Card
- Cash
- (Future) Crypto, online wallet, Apple Pay, etc.

Hardcoding:

```go
if method == "card" { ... }
```

everywhere is bad practice.

Factory Method centralizes creation logic.

---

### **Code: `internal/creational/factory/payment_factory.go`**

```go
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
```

### **Usage:**

```go
payment := factory.GetPaymentMethod("card")
payment.Pay(120)
```

### **Why Factory Method?**
- Client code does not know or care which concrete class is used
- Extending with new methods requires no modification of client code
- Enforces polymorphic behavior

---

#  **5. Client Code Demonstration**

### **File: `cmd/client/main.go`**

```go
package main

import (
    "fmt"
    "example.com/creational/internal/creational/singleton"
    "example.com/creational/internal/creational/builder"
    "example.com/creational/internal/creational/factory"
)

func main() {
    // Singleton usage
    config := singleton.GetConfig()
    fmt.Println("Config value:", config.Value)

    // Builder usage
    bouquet := builder.NewBouquetBuilder().
        SetName("Spring Mix").
        AddFlower("Rose").
        AddFlower("Tulip").
        Build()
    fmt.Println("Bouquet:", bouquet.Name, bouquet.Flowers)

    // Factory Method usage
    payment := factory.GetPaymentMethod("card")
    payment.Pay(120)
}
```

---

#  **6. Project Structure**

```
cmd/client/main.go
internal/domain/bouquet.go
internal/creational/
  singleton/
    config.go
  builder/
    bouquet_builder.go
  factory/
    payment_factory.go
README.md
```

This structure fully aligns with the recommended TMPS architecture:
- **client**
- **domain**
- **factory**
- **models**

---

# 📈 **7. Results & Output Example**

Running:

```bash
go run ./cmd/client
```

Produces:

```
Config value: Bloomify Default Config
Bouquet: Spring Mix [Rose Tulip]
Paid with card: 120
```

---

# 📝 **8. Conclusions**

Through this laboratory work:

- **Singleton** ensured consistent configuration across the project
- **Builder** simplified creation of complex bouquet objects
- **Factory Method** abstracted payment method selection, ensuring extensibility

The project meets all TMPS requirements:

✔ 3 Creational Patterns implemented  
✔ Object instantiation abstracted and optimized  
✔ Clear modular architecture  
✔ Fully documented  
✔ No frameworks or external libraries

Bloomify now has a maintainable and scalable foundation that can easily evolve into a full e-commerce platform.

---


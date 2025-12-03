# 🧩 Laboratory Work #1 – SOLID Principles

### 🎯 Objective
Implement and understand **three SOLID principles** in a language of my choice (Go):
1. **SRP – Single Responsibility**
2. **OCP – Open/Closed**
3. **DIP – Dependency Inversion**

---

## 📘 1️⃣ Single Responsibility Principle (SRP)

> A class should have only one reason to change.

**Example:**  
`Report` only handles storing and displaying text.  
It does **not** manage payments, calculations, or notifications.

```go
report := Report{"Lab 1 - SOLID", "Understanding SRP, OCP, and DIP"}
report.Display()
```

✅ Keeps code focused, testable, and easy to maintain.

---

## 📗 2️⃣ Open/Closed Principle (OCP)

> Software entities should be open for extension but closed for modification.

**Example:**  
`Shape` is an interface with `Area()`.  
We can extend functionality by adding new shapes without editing existing code.

```go
circle := Circle{Radius: 3}
square := Square{Side: 4}
PrintArea(circle)
PrintArea(square)
```

✅ Encourages flexibility and scalability.

---

## 📙 3️⃣ Dependency Inversion Principle (DIP)

> High-level modules should depend on abstractions, not on concrete implementations.

**Example:**  
`Notification` depends on the `MessageSender` interface.  
You can plug in `EmailSender`, `SMSSender`, or any new type that implements `Send()`.

```go
notifyEmail := Notification{Sender: EmailSender{}}
notifySMS := Notification{Sender: SMSSender{}}
notifyEmail.Alert()
notifySMS.Alert()
```

✅ Allows swapping modules without breaking core logic.

---

## ▶️ How to Run

```bash
go run .
```

Then input:

```
1
```

**Expected Output:**
```
SRP → Report manages only text
OCP → Shapes extend without modifying logic
DIP → Notification depends on abstractions
```
---

## 🧾 Summary

| Principle | Meaning | Example |
|------------|----------|----------|
| **SRP** | One class = one responsibility | `Report` |
| **OCP** | Extend without modifying | `Shape` interface |
| **DIP** | Depend on abstractions | `MessageSender` + `Notification` |

---

👤 Developed by **Bujor Alexandru**, Group **FAF-231**

 
 

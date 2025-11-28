package strategy

import "fmt"

type DeliveryStrategy interface { Deliver(orderID int) }

type Courier struct{}
func (c *Courier) Deliver(id int) { fmt.Println("Courier delivers", id) }

type Drone struct{}
func (d *Drone) Deliver(id int) { fmt.Println("Drone delivers", id) }

type Context struct { Strategy DeliveryStrategy }

func (c *Context) SetStrategy(s DeliveryStrategy) { c.Strategy = s }
func (c *Context) Execute(id int) { c.Strategy.Deliver(id) }

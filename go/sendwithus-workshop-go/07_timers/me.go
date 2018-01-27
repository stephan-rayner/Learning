package main

import (
    "fmt"
    "time"
)

type Computer struct {
	Brand string
	Model string
	Price int
}

func (c *Computer) Describe() {
    fmt.Printf("%s %s $%d\n", c.Brand, c.Model, c.Price)
}

func (c *Computer) StartTimer(t time.Duration) {
    fmt.Println("Starting Timer for", t)
    time.Sleep(t)
    fmt.Println("End of Timer")
}

func main() {
    computer := Computer{
        Brand: "Apple",
        Model: "Macbook",
        Price: 1000,
    }
    
    t := 3 * time.Second
    computer.StartTimer(t)
}
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
    fmt.Println("Starting timer for", t)
    time.Sleep(t)
    fmt.Println("Time up after", t)
}

func main() {
    computer := Computer{
        Brand: "Apple",
        Model: "Macbook",
        Price: 1000,
    }
    
	go computer.StartTimer(5 * time.Second)
	go computer.StartTimer(2 * time.Second)
	go computer.StartTimer(1 * time.Second)	
	go computer.StartTimer(4 * time.Second)
    
    // IGNORE THIS:
    // This is a hack, so the program doesn't quit before the goroutine finishes
    time.Sleep(10 * time.Second)
}
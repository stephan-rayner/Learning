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

func (c *Computer) StartTimer(channel chan string, t time.Duration) {
    fmt.Println("Starting timer...")
    time.Sleep(t)
    
	// TODO: Push the "Time up!" string to the channel
	channel <- "Times up!"
}

func main() {
    computer := Computer{
        Brand: "Apple",
        Model: "Macbook",
        Price: 1000,
    }
    
    channel := make(chan string)
    
    t := 3 * time.Second
    go computer.StartTimer(channel, t)
    
    select {
        case msg := <-channel:
            fmt.Println(msg)
    }

    fmt.Println("Exited")
}

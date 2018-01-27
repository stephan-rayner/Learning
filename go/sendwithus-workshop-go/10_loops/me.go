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
    channel <- "Time up!"
}

func (c *Computer) StartInterval(channel chan string, t time.Duration) {
	
	defer close(channel) 
	fmt.Println("Starting interval...")

    // TODO: Use a "for" loop to push a message to the channel every t seconds
    for i := 0; i < 10; i++ {
		time.Sleep(t)
		channel <- fmt.Sprintf("Interval %d up!", i)
	} 
	
    // TODO: Close the channel after the loop finishes
}

func main() {
    computer := Computer{
        Brand: "Apple",
        Model: "Macbook",
        Price: 1000,
    }
    
    channel := make(chan string)
    
    t := 1 * time.Second
    go computer.StartInterval(channel, t)

    computer.Describe()
    
    // TODO: Use a "for" loop to watch for messages on the channel
    
    fmt.Println("Exited")
}
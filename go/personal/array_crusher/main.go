package main

import "fmt"
import "time"


func processStuff(value int, channel chan<- int) {
    time.Sleep(time.Second * 1)
    channel <- value
}


func main() {
    startTime := time.Now()

    numSeconds := 86400 // Number of seconds in a day.
    channel := make(chan int, numSeconds)

    for i := 0; i < numSeconds; i++ { 
        go processStuff(i, channel)
    }

    for i := 0; i < numSeconds; i++ {
        fmt.Println(<-channel)
    }

    fmt.Println(time.Now().Sub(startTime))
}

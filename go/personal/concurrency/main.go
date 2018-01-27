package main

import "fmt"

func inc(i int, c chan int) {
    c <- i + 1
}

func learnConcurrency() {
    c := make(chan int)
    go inc(0, c)
    go inc(10, c)
    go inc(-805, c)

    a := <-c
    fmt.Println(a, <-c, <-c)
}

func main () {
    learnConcurrency()
}

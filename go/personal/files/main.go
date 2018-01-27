package main

import "os"
import "fmt"
import "log"

func create() {
    println("creating stuff")
    file, _ := os.Create("cheese.txt")
    defer file.Close()
    fmt.Fprint(file, "CHEESE BITCHES :P")
}
/*
func write() {
    file, err := os.Open("cheese.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    fmt.Fprint(file, "MORE CHEESE BITCHES :)")
    file.Write("MORE CHEESE BITCHES :)")
}
*/
func read() {
    println("reading stuff")
    file, err := os.Open("cheese.txt")
    if err != nil {
        log.Fatal(err)
    }
    
    defer file.Close()
    data := make([]byte, 100)
    stuff, _ := file.Read(data)
    fmt.Printf("read %d bytes: %q\n", stuff, data[:stuff])
}

func main() {
    create()
    //write()
    read()
}

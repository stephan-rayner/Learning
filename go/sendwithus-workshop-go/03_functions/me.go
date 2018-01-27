package main

import "fmt"

func print(textInput ...string) {
    output := ""
    for _, text := range textInput {
        output = output + " " + text
    }
    fmt.Println(output)
}

func main() {
    print("Hello", "World")
}
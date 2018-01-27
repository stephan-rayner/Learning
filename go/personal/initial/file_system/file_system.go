package file_system

import "os"
import "fmt"
import "log"

func Create(file_name string) {
    println("creating stuff")
    file, _ := os.Create(file_name)
    defer file.Close()
    fmt.Fprint(file, "CHEESE BITCHES :P")
}

func Write(file_name string, stuff_to_write string) {
    println("writing stuff")
    file, err := os.Open(file_name)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    fmt.Fprint(file, "MORE CHEESE BITCHES :)")
    file.Write([]byte(stuff_to_write))
}

func Read(file_name string) {
    println("reading stuff")
    file, err := os.Open(file_name)
    if err != nil {
        log.Fatal(err)
    }
    
    defer file.Close()
    data := make([]byte, 100)
    stuff, _ := file.Read(data)
    fmt.Printf("read %d bytes: %q\n", stuff, data[:stuff])
}

package main

import (
    "fmt"
    fs "./file_system"
    "net/http"
    "ioutil"
)


func main() {
    rs, _ := http.Get("https://jsonplaceholder.typicode.com/posts")
    defer rs.Body.Close()

    bodyBytes, err := ioutil.ReadAll(rs.Body)
    if err != nil {
        panic(err)
    }
 
    bodyString := string(bodyBytes)
    
    fmt.Println(bodyString)
}

func old(){
    // floyd.Triangle(4)
    // fmt.Println("============================")
    // number := 102355529302
    
    // var ans string

    // if checks.Is_even(number) {
    //     ans = "Yes"
    // } else {
    //     ans = "No"
    // }
    // fmt.Println("Is", number, "even?")
    // fmt.Println(ans)
    
    // Swtich cases don't fall through.
    x := 2
    switch x {
    case 0:
        println("zero")
    case 1:
        println("one")
    case 2:
        println("two")
    case 3:
        println("three")
    default:
        println("I just don't even know")
    }

    //a := make(chan int)
    var m = map[string]int{"one":1, "two":2, "three":3}
    fmt.Println("Using a map", m["two"])   
    println("CHEESE")

    fs.Create("./data/txt.txt")
    fs.Write("./data/txt.txt", "Fuck ya cheese!")
    fs.Read("./data/txt.txt")
    // "MORE CHEESE BITCHES :)"
}


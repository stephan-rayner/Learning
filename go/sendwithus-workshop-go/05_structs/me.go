package main

import "fmt"

type Computer struct {
    Brand string
    Model string
    Price int
}

func main() {
    c1 := Computer{
        Brand: "Apple",
        Model: "Macbook",
        Price: 1000,
    }

	c2 := Computer{
		Brand: "thing",
		Model: "thang",
		Price: 500,
	}

    fmt.Println(c1)
    fmt.Println(c2)
}
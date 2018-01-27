package floyd

import "fmt"

func Triangle(line_count int) {
    number := 1
    for i := 1; i <= line_count; i++ {
        for j := 1; j <= i; j++ {
            fmt.Print(" ", number, " ")
            number ++
        }
        fmt.Println("")
    }
}
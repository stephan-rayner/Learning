package main


func learnNamedReturns(x, y int) (z int) {
    z = x * y
    return
}

func main() {
    result := learnNamedReturns(2, 5)
    println(result)
}

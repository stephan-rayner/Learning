package main

func main() {
	for x := 0; x < 100; x++ {

		if x%3 == 0 && x%5 == 0 {
			println("FizzBuzz")
		} else if x%3 == 0 {
			println("Fizz")
		} else if x%5 == 0 {
			println("Buzz")
		} else {
			println(x)
		}
	}

}

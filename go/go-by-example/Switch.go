package main

import "time"

func whatIsThis(i interface{}) {
	switch t := i.(type) {
	case bool:
		println("Boolean")
	case int:
		println("Integer")
	default:
		println("Implement Types %T\n", t)
	}
}

func main() {
	i := 42

	switch i {
	case 1:
		println("ONE")
	case 2:
		println("TWO")
	case 3:
		println("THREE")
	case 42:
		println("Life the universe and everything")
	default:
		println("Beats the heck out of me.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		println("Good Morning")
	case t.Hour() >= 12 && t.Hour() < 16:
		println("Good Afternoon")
	case t.Hour() >= 16 && t.Hour() < 20:
		println("Good Evening")
	case t.Hour() >= 20:
		println("Good Night")
	default:
		println("I don't know that time of day it is, I am going to bed.")
	}

	whatIsThis(true)
	whatIsThis(1)
	whatIsThis(2.2)
	whatIsThis("Cheese")
}

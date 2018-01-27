package main

import (
	"math"
)

func main() {
	const n = 500000000
	const d = 3e20 / n
	println(d)
	println(int64(d))
	println(math.Sin(180))
}

package main

func main() {
	i := 1
	for i <= 3 {
		println(i)
		i = i + 1
	}

	println("================")

	for j := 0; j < 10; j++ {
		println(j)
	}

	println("================")

	for {
		println("loop")
		break
	}

	println("================")

	for k := 0; k <= 8; k++ {
		if k%2 != 0 {
			continue
		}
		println("Even Number:", k)
	}

}

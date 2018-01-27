package main

import (
	"fmt"
)

func main() {
	a := car{
		gasPedal:      65000,
		breakPedal:    0,
		steeringWheel: 12561,
		topSpeedKph:   225.0}
	fmt.Println(a.kph())
	fmt.Println(a.mph())
	a.setTopSpeed(500)
	fmt.Println(a.kph())
	fmt.Println(a.mph())
}

package main

const u16bitmax float64 = 65533
const kphMultiple float64 = 1.60934

// This is my car
type Car struct {
	gasPedal      uint16 // 0 -> 65535
	breakPedal    uint16
	steeringWheel int16
	topSpeedKph   float64
}

func (c Car) kph() float64 {
	return float64(c.gasPedal) * (c.topSpeedKph / u16bitmax)
}

func (c Car) mph() float64 {
	return (float64(c.gasPedal) * (c.topSpeedKph / u16bitmax)) / kphMultiple
}

func (c *Car) setTopSpeed(speed float64) {
	c.topSpeedKph = speed
}

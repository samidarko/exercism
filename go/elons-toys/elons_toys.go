package elon

import "fmt"

// Drive the car
func (c *Car) Drive() {
	batteryLeft := c.battery - c.batteryDrain
	if batteryLeft >= 0 {
		c.battery = batteryLeft
		c.distance += c.speed
	}
}

// DisplayDistance display distance
func (c Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// DisplayBattery display battery
func (c Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

// CanFinish return true if car can finish
func (c *Car) CanFinish(trackDistance int) bool {
	accelerations := trackDistance / c.speed
	return c.battery-c.batteryDrain*accelerations >= 0
}

package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmhMultiple float64 = 1.60934

type car struct {
	gasPedal      uint16 //min: 0,      max: 65535
	brakePedal    uint16 //min: 0,      max: 65535
	steeringWheel int16  //min: -32768  max: 32768
	topSpeedKmh   float64
}

/**
kmh() and mph() are value functions, meaning that they receive a copy of the struct rather than the struct itself.
Modifying the values in the function therefore only modifies the copy's value and not the originating struct instance that was
passed in to the function
*/
// Here the () brackets that come before the function name associate the function with the struct "car"
func (c car) kmh() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh / usixteenbitmax)
}

func (c car) mph() float64 {
	// c.topSpeedKmh = 25 // this modifies the instance in this functions context
	return float64(c.gasPedal) * (c.topSpeedKmh / usixteenbitmax / kmhMultiple)
}

func (c *car) pointerMph() float64 {
	c.topSpeedKmh = 25 // this modifies the original instance of the struct that was passed into the function
	return float64(c.gasPedal) * (c.topSpeedKmh / usixteenbitmax / kmhMultiple)
}

// This is a setter function: use the "*" to specify a change to the instance of the struct
func (c *car) superCharge() {
	var newSpeed = c.topSpeedKmh + 100.0
	c.topSpeedKmh = newSpeed
}

// Value function
func newTopSpeed(c car, newSpeed float64) {
	c.topSpeedKmh = newSpeed
	fmt.Println("Car's top speed is now: ", c.topSpeedKmh)
}

func (c *car) newTopSpeedPointer(newSpeed float64) {
	fmt.Println("new top speed pointer called...")
	c.topSpeedKmh = newSpeed
}

func main() {

	aCar := car{65524, 0, 12562, 225.0}

	fmt.Println("gas_pedal:", aCar.gasPedal)
	fmt.Println("speed (km/h):", aCar.kmh())
	fmt.Println("speed (mph):", aCar.mph())
	aCar.superCharge()
	fmt.Println("new speeds (super-charged)")
	fmt.Println("speed (km/h):", aCar.kmh())
	fmt.Println("speed (mph):", aCar.mph())
	fmt.Println("using the pointer function...")
	fmt.Println("speed (mph):", aCar.pointerMph())
	fmt.Println("the new top speed (km/h) is ", aCar.topSpeedKmh)
	newTopSpeed(aCar, 250.0)
	fmt.Println("the new top speed (km/h) is ", aCar.topSpeedKmh)
	aCar.newTopSpeedPointer(100.25)
	fmt.Println("the new top speed (km/h) is ", aCar.topSpeedKmh)
}

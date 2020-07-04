package main

import "fmt"

type Liters float64
type Milliliters float64
type Gallons float64

func main() {
	soda := Liters(2)
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.ToGallons())
	water := Milliliters(500)
	fmt.Printf("%0.3f milliliters equals %0.3f gallons\n", water, water.ToGallons())
}
func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

// this is a function, not a method
func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

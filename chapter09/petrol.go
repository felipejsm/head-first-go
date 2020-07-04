package main

import "fmt"

type Liters float64
type Gallons float64

func main() {
	var carFuel Gallons
	var busFuel Liters
	carFuel = Gallons(10.0)
	busFuel = Liters(240.0)
	fmt.Println(carFuel, busFuel)
	fmt.Println(convertGalllonsToLiters(carFuel))
	fmt.Println(convertLitersToGallons(busFuel))
}
func convertLitersToGallons(busFuel Liters) Gallons {
	return Gallons(busFuel * 0.264)
}
func convertGalllonsToLiters(carfuel Gallons) Liters {
	return Liters(carfuel * 3.785)
}

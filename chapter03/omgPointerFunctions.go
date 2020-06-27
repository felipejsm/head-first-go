package main

import "fmt"

func main() {
	var myFloatPointer *float64 = createPointer()
	fmt.Println(*myFloatPointer)
	myLight := true
	fmt.Printf("My light is %t\n", myLight)
	onOrOff(&myLight)
	fmt.Printf("Now my light is %t\n", myLight)
}

func createPointer() *float64 {
	var myFloat = 98.5
	return &myFloat
}

func onOrOff(myLight *bool) {
	*myLight = !(*myLight)

}

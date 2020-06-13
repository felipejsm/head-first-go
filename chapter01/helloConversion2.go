package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInt int = 3
	fmt.Println("Convert int to float64 ", float64(myInt))
	fmt.Println(reflect.TypeOf(myInt))
	fmt.Println(reflect.TypeOf(float64(myInt)))
}

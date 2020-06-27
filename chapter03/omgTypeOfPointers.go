package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInt int
	var myFloat float64
	var myBool bool
	myBoolPointer := &myBool
	fmt.Println(reflect.TypeOf(&myInt))
	fmt.Println(reflect.TypeOf(&myFloat))
	fmt.Println(reflect.TypeOf(&myBoolPointer))
	fmt.Println("---------------------")
	myInt = 64
	myIntPointer := &myInt
	myFloat = 3.1415
	myFloatPointer := &myFloat
	myBool = true
	fmt.Println(*myIntPointer)
	fmt.Println(*myFloatPointer)
	fmt.Println(*myBoolPointer)
	*myIntPointer = 46
	fmt.Println(*myIntPointer)
	fmt.Println(myInt)
}

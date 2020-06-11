package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(2020))
	fmt.Println(reflect.TypeOf("Lil Pump"))
	fmt.Println(reflect.TypeOf(3.15))
	fmt.Println(reflect.TypeOf('Ç'))
	fmt.Println(reflect.TypeOf(false))
}

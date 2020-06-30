package main

import "fmt"

func main() {
	myArray := [5]string{"a", "b", "c", "d", "e"}
	i, j := 1, 4
	myInnerArray := myArray[i:j]
	fmt.Println(myInnerArray)
}

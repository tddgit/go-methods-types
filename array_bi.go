package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	var arr1 [5]int
	arr2 := [5]int{10, 20, 30, 40, 50}
	arr3 := [...]int{10, 20, 30, 40, 50}
	arr4 := [5]int{1: 10, 2: 20}
	arr4[3] = 43
	arr4[1] = 56

	arr5 := [5]*int{0: new(int), 1: new(int)}
	*arr5[0] = 10
	*arr5[1] = 20

	fmt.Println(arr1, arr2, arr3, arr4, arr5)
}

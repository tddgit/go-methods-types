package main

import (
	"fmt"
	"unsafe"
)

func main() {

	array := [...]int{0, 1, -2, 3, 4}
	pointer := &array[0]
	fmt.Print(*pointer, " ")

	memoryAddress := uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])

	for i := 0; i < len(array)-1; i++ {
		pointer =
	}

	var value int64 = 5
	var p1 = &value
	var p2 = (*int32)(unsafe.Pointer(p1))

	fmt.Println("*p1:  ", *p1)
	fmt.Println("*p2:  ", *p2)
	*p1 = 5432366955555555669
	fmt.Println(value)
	fmt.Println("*p2:  ", *p2)
	*p1 = 546789
	fmt.Println(value)
	fmt.Println("*p2:  ", *p2)
}

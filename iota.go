package main

import "fmt"

type Digit int
type Power2 int

const PI = 3.1415996

const (
	C1 = "C1C1C1"
	C2 = "C2C2C2"
	C3 = "C3C3C3"
)
const (
	Zero Digit = iota
	One
	Two
	Three
	Four
)
const (
	p2_0 Power2 = 1 << iota
	_
	p2_2
	_
	p2_4
	_
)

func main() {
	fmt.Println("======================================================")

	anArray := [5]int{0, 1, -1, 2, -2}
	for i, value := range anArray {
		fmt.Println("index", i, "value", value)
	}

	fmt.Println("BASIC LOOP")
	for i := 0; i < 100; i++ {
		if i%20 == 0 {
			fmt.Println("i%20 == 0")
			continue
		}
		if i == 95 {
			fmt.Println("i == 95")
			break
		}
	}

	fmt.Println("WHILE LOOP")
	i := 10
	for {
		if i < 0 {
			break
		}
		fmt.Print(i, " ")
		i--
	}

	fmt.Println("\nDO WHILE LOOP")

	i = 10
	anExpression := true
	for ok := true; ok; ok = anExpression {
		if i > 10 {
			anExpression = false
		}
		fmt.Print(i, " ")
		fmt.Print("ok -> ", ok)
		i++
	}

	fmt.Println()

	c1 := 12 + 1i
	c2 := complex(5, 7)
	var c3 complex64 = complex64(c1 + c2)

	cZero := c3 - c3
	fmt.Printf("cZero: %T: %v", cZero)

	aSliceLiteral := []int{1, 2, 3, 4, 5}
	integer := make([]int, 20)

	for i := 0; i < len(integer); i++ {
		integer[i] = i
		fmt.Println(integer[i])
	}

	aSliceLiteral = nil
	integer = nil

	integer = append(integer, 12345)

	fmt.Println(aSliceLiteral, integer)

	const s1 = 1234
	var v1 float32 = s1 * 12
	fmt.Println(v1)
	fmt.Println(PI)

	fmt.Println(One)
	fmt.Println(Two)

	fmt.Println("2^0", p2_0)
	fmt.Println("2^2", p2_2)
	fmt.Println("2^4", p2_4)
}

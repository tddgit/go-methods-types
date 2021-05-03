package main

import "fmt"

type twoInts struct {
	X int64
	Y int64
}

func regularFunction(a, b twoInts) twoInts {
	temp := twoInts{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
	return temp
}

func (a twoInts) method(b twoInts) twoInts {
	temp := twoInts{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
	return temp
}

type Reader1 interface {
	Read(p []byte) (n int, err error)
}

type Writer1 interface {
	Write(p []byte) (n int, err error)
}

type a struct {
	XX int
	YY int
}

type b struct {
	AA string
	XX int
}

type c struct {
	A a
	B b
}

func (A a) A() {
	fmt.Println("Function A() for A")
}

func (B b) A() {
	fmt.Println("Function A() for B")
}

type first struct{}

func (a first) F() {
	a.shared()
}
func (a first) shared() {
	fmt.Println("This is shared() from first!")
}

type second struct {
	first
}

func (a second) shared() {
	fmt.Println("This is shared() from second!")
}

func main() {
	first{}.F()
	second{}.shared()
	second{}.first.shared()
	l := second{}
	j := l.first
	j.F()

	var k c
	k.A.A()
	k.B.A()

	i := twoInts{
		X: 1,
		Y: 2,
	}
	j := twoInts{
		X: 3,
		Y: 4,
	}
	fmt.Println(regularFunction(i, j))
	fmt.Println(i.method(j))
	fmt.Println()
}

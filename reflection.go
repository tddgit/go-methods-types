package main

import (
	"fmt"
	"log"
	"log/syslog"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type square struct {
	X float64
}

func (s square) Area() float64 {
	return s.X * s.X
}
func (s square) Perimeter() float64 {
	return s.X * 4
}

type circle struct {
	R float64
}

func (c circle) Area() float64 {
	return c.R * c.R
}
func (c circle) Perimeter() float64 {
	return 2 * (math.Pi * c.R)
}

func Calculate(x Shape) {
	_, ok := x.(circle)
	if ok {
		fmt.Println("It is a circle")
	}

	v, ok := x.(square)
	if ok {
		fmt.Println("It is a square", v)
	}

	fmt.Println(x.Area())
	fmt.Println(x.Perimeter())
}

type rect struct {
	a, b float64
}

func tellInterface(x interface{}) {
	switch v := x.(type) {
	case square:
		fmt.Println("This is a square")
	case circle:
		fmt.Printf("%v is a circle\n", v)
	case rect:
		fmt.Printf("This is %T\n", v)
	default:
		fmt.Printf("Unrecognized shape %T\n", v)
	}
}

type a struct {
	X int
	Y float64
	Z string
}

type b struct {
	F int
	G int
	H string
	I float64
}

func main() {
	cmd := exec.Command("cls", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println("\033[2J")
	fmt.Println("=================================================")

	programName := filepath.Base(os.Args[0])
	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)

	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}
	log.Println("LOG_INFO _ LOG_LOCAL7: LOGGING in Go!")

	sysLog, err = syslog.New(syslog.LOG_MAIL, "Some program!")

	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}

	log.Println("LOG_MAIL: Logging in Go!")
	fmt.Println("Will you see this?")

	x0 := 1000
	xRef1 := reflect.ValueOf(&x0).Elem()
	fmt.Println(xRef1, reflect.TypeOf(xRef1))
	xType := xRef1.Type()
	fmt.Println(xType, reflect.TypeOf(xType))

	x := square{10}
	fmt.Println("Perimeter square: ", x.Perimeter())
	fmt.Println("Area square: ", x.Area())

	Calculate(x)
	y := circle{10}
	fmt.Println("Perimeter circle: ", y.Perimeter())
	fmt.Println("Area circle:", y.Area())

	Calculate(y)

	z := rect{10, 5}
	tellInterface(x)
	tellInterface(y)
	tellInterface(z)

	var myInt interface{} = 123

	k, ok := myInt.(int)

	if ok {
		fmt.Println("Success: ", k)
	} else {
		fmt.Println("Failed without panicking")
	}

	i := myInt.(int)
	fmt.Println(i)

	j, ok := myInt.(bool)
	if ok {
		fmt.Println(j)
	} else {
		fmt.Println("Failed without panicking")
	}

}

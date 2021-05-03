package main

import (
	"fmt"
	"log"
	"os"
)

func d1() {
	for i := 3; i > 0; i-- {
		defer fmt.Print(i)
	}
}

func d2() {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
}

func d3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
}

func one(aLog *log.Logger) {
	aLog.Println("--- FUNCTION one ---------------")
	defer aLog.Println("----- FUNCTION one deferred -----------------")

	for i := 0; i < 10; i++ {
		aLog.Print(i, " ")
	}
	fmt.Println()
}

func two(aLog *log.Logger) {
	aLog.Println("--------------FUNCTION two---------------")
	defer aLog.Println("------------------Function two defered---------------")

	for i := 10; i > 0; i-- {
		aLog.Println(i)
	}
}

func a() {
	fmt.Println("Inside a()")
	defer func() {
		if c := recover(); c != nil {
			fmt.Println("Recover inside a()")
		}
	}
}

func main() {
	LOGFILE := "log.txt"
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)

	}
	defer f.Close()

	iLog := log.New(f, "logDefer     ", log.LstdFlags)
	iLog.Println("Hello there!")
	iLog.Println("Another log entry")

	one(iLog)
	two(iLog)

	fmt.Println("===========================================")
	d1()
	fmt.Println()
	d2()
	fmt.Println()
	d3()
	fmt.Println()

}

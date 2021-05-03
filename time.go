package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("======================================================")

	fmt.Println("Epoch time:", time.Now().Unix())
	t := time.Now()
	fmt.Println(t)

	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(t.Weekday(), t.Day(), t.Month(), t.Year())

	time.Sleep(time.Second)
	time.Sleep(time.Minute / 10)
	t1 := time.Now()

	fmt.Println("Time difference: ", t1.Sub(t))

	formatT := t.Format("01 January 2006")
	fmt.Println(formatT)
	loc, err := time.LoadLocation("Europa/Paris")
	moscowTime := t.In(loc)
	fmt.Println("Moscow", moscowTime)

	parsedTime, err := time.Parse("01 January 2006", "01 January 2006")

	if err != nil {
		fmt.Println("Not a time string")
	}

	fmt.Println(parsedTime.Format("01 January 2006"))

}

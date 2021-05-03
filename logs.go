package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"log/syslog"
	"os"
	"strconv"
)

func main() {
	
	if len(os.Args) == 1 {
		fmt.Println("Please give one more floats!")
		os.Exit(1)
	}
	arguments := os.Args
	min, _ := strconv.ParseFloat(arguments[1], 64)
	max, _ := strconv.ParseFloat(arguments[1], 64)

	for i := 2; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)

	var f *os.File
	f = os.Stdin
	defer (func() {
		err := f.Close()

		if err != nil {
			log.Fatal(err)
		}
	})()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}

	myString := ""
	arguments = os.Args

	if len(arguments) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = arguments[1]
	}

	_, err1 := io.WriteString(os.Stdout, myString)
	_, err2 := io.WriteString(os.Stdout, "\n")

	if err1 != nil || err2 != nil {
		log.Fatal(err1, err2)
	}

	sysLog, err := syslog.New(syslog.LOG_ALERT|syslog.LOG_MAIL, "Some program!")
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}

}

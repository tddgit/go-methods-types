package main

import (
	"fmt"
	"time"
)

func main() {

	jobs := make(chan Job)
	done := make(chan bool, len(jobList))

	go func() {
		for _, job := range jobList {
			jobs <- job
		}
		close(jobs)
	}

	go func() {
		for job := range jobs {
			fmt.Println(job)
			done <- true
		}
	}

	for i := 0; i < len(jobList); i++ {
		<-done
	}
	fmt.Println("=======================")

	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepyGopher1(i, c)
	}
	for i := 0; i < 5; i++ {
		gotherID := <-c
		fmt.Println("gother", gotherID, "has finished sleeping")
	}

}

func sleepyGopher1(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("...", id, "snore ...")
	c <- id
}

func startGoroutines() {
	go sleepyGopher()
	fmt.Println("main func")
	time.Sleep(4 * time.Second)

	for i := 0; i < 5; i++ {
		go sleepyGopher()
	}
	fmt.Println("Goroutines cycle")
	time.Sleep(4 * time.Second)
}

func sleepyGopher() {
	time.Sleep(3 * time.Second)
	fmt.Println("... snore ...")

}

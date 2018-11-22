package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// You can ignore the rest, it's just a test.

func workertest() {

	// Jobs
	ss := [...]string{
		"This",
		"Is",
		"Test",
		"Input",
		"To",
		"Simulate",
		"A",
		"Workload",
	}

	// Channels
	fmt.Println("Creating channels.")
	jobs := make(chan string, runtime.NumCPU())
	results := make(chan time.Duration)
	chanID := make(chan int)

	// Workers
	fmt.Println("Creating workers.")
	for w := 1; w <= runtime.NumCPU(); w++ {
		go testworker(w, jobs, results, chanID)
	}

	// Jobs
	fmt.Println("Giving jobs.")
	for _, s := range ss {
		jobs <- s
	}
	close(jobs)

	// Results
	for i := 0; i < len(ss); i++ {
		result := <-results
		fmt.Println("Worker", <-chanID, "|", "Duration", result, "|", i+1, "/", len(ss))
	}
}

func testworker(id int, inputs <-chan string, outputDuration chan<- time.Duration, outputID chan<- int) {
	for _ = range inputs {
		begin := time.Now()
		time.Sleep(time.Duration(float64(time.Minute) * rand.New(rand.NewSource(time.Now().UnixNano())).Float64()))
		outputDuration <- time.Since(begin)
		outputID <- id
	}
}

/* Implements mutual exclusion with sync/mutex */

package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

const (
	MAX_COUNT  = 10000000
	GOROUTINES = 2
)

var (
	counter = 0
	mutex int32 = 0
)

func lock() {
	for ! atomic.CompareAndSwapInt32(&mutex, 0, 1) {}
}

func unlock() {
	mutex = 0
}

func run(id, counts int, done chan bool) {
	for i := 0; i < counts; i++ {
		lock()
		counter++
		unlock()
	}
	fmt.Printf("End %d counter: %d\n", id, counter)
	done <- true
}

func main() {
	runtime.GOMAXPROCS(GOROUTINES)
	done := make(chan bool, 1)

	for i := 0; i < GOROUTINES; i++ {
		go run(i, MAX_COUNT/GOROUTINES, done)
	}

	for i := 0; i < GOROUTINES; i++ {
		<-done
	}

	fmt.Printf("Counter value: %d Expected: %d\n", counter, MAX_COUNT);
}

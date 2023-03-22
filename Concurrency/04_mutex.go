package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	const c = 100000
	shared := 1

	//set to number of cores available
	fmt.Println("GOMAXPROCS ", runtime.GOMAXPROCS)
	fmt.Println("Num of goroutines START : ", runtime.NumGoroutine())

	//spun 100 go routines
	wg.Add(c)
	for x := 1; x <= c; x++ {
		fmt.Println("I am GO routine & shared value is ", x, shared)
		go func() {
			mu.Lock()
			v := shared
			runtime.Gosched()

			//runtime.Gosched() basically says go ahead and run another goroutine if you want.
			//In other words, it yields the processor, allowing other goroutines to run.
			//It increases the randomness so that the probability of goroutines interfering with each other and
			//reading-writing in different order increases.
			v++
			shared = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Go routine", x, "is over DONE")

	}
	fmt.Println("Num of goroutines END : ", runtime.NumGoroutine())
	fmt.Println("Shared value : ", shared)
	wg.Wait()
	fmt.Println("waiting is over")
}

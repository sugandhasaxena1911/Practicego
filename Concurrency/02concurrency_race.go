package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	/*
			run this program without waitgroup , notice that main gets over without all the spunned go routines
			being processed (ie without printing  :
			"Inside go routine NO 100


		with wg :
		Inside go routine NO 100 is printed because main will wait till all spunned routines are processed

			)
	*/

	const c = 100
	shared := 1

	//set to number of cores available
	fmt.Println("GOMAXPROCS ", runtime.GOMAXPROCS)
	fmt.Println("Num of goroutines START : ", runtime.NumGoroutine())

	//spun  go routines
	wg.Add(c)
	for x := 1; x <= c; x++ {
		fmt.Println("I am GO routine & shared value is ", x, shared)
		go func() {
			//Using lock will make sure that only one outine has access to shared variable at one time .
			//Therefore you will not find any case when below :
			//"Inside go routine processing & BEFORE shared val is"
			//is printed consecutively .
			//always BEFORE and AFTER are prnted and then next BEFORE and AFTER
			mu.Lock()
			v := shared
			fmt.Println("Inside go routine processing & BEFORE shared val is   ", shared)
			//runtime.Gosched()
			//runtime.Gosched() basically says go ahead and run another goroutine if you want.
			//In other words, it yields the processor, allowing other goroutines to run.
			//It increases the randomness so that the probability of goroutines interfering with each other and
			//reading-writing in different order increases.
			v++
			shared = v
			fmt.Println("Inside go routine processing & AFTER shared val is   ", shared)

			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Go routine", x, "is now spunned")

	}
	fmt.Println("Num of goroutines END : ", runtime.NumGoroutine())
	fmt.Println("Shared value : ", shared)
	wg.Wait()
	fmt.Println("waiting is over")
}

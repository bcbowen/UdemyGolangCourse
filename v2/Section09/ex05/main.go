package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Hands-on exercise #5
Fix the race condition you created in exercise #4 by using package atomic


*/
var wg sync.WaitGroup

func main() {
	fmt.Println("Hello, playground")
	var counter int64
	wg.Add(99)
	for i := 0; i < 99; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("The counter is: ", counter)
}

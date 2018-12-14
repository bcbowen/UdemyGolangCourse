package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
Using goroutines, create an incrementer program
have a variable to hold the incrementer value
launch a bunch of goroutines
each goroutine should
read the incrementer value
store it in a new variable
yield the processor with runtime.Gosched()
increment the new variable
write the value in the new variable back to the incrementer variable
use waitgroups to wait for all of your goroutines to finish
the above will create a race condition.
Prove that it is a race condition by using the -race flag

*/
var wg sync.WaitGroup

func main() {
	fmt.Println("Hello, playground")
	counter := 0
	var mutex sync.Mutex
	wg.Add(99)
	for i := 0; i < 99; i++ {
		go func() {
			mutex.Lock()
			val := counter
			runtime.Gosched()
			val++
			counter = val
			mutex.Unlock()
			wg.Done()

		}()
	}
	wg.Wait()
	fmt.Println("The counter is: ", counter)
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	c := make(chan int, 100)
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(start int) {
			for j := start; j < start+10; j++ {
				c <- j
			}
			wg.Done()
		}(i * 10)
	}
	wg.Wait()
	close(c)

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("Ex #7 done!")
}

package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)
	c <- 23
	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}

package main

import (
	"fmt"
)

type customErr struct {
	message string
}

func (ce customErr) Error() string {
	return ce.message
}

func main() {
	c1 := customErr{
		message: "oh fuckballs!",
	}
	foo(c1)
}

func foo(e error) {
	fmt.Println(e)
}

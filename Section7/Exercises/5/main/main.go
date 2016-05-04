package main

import "fmt"

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}

func foo(args ...int) {
	for i, a := range args {
		fmt.Printf("%v. %v\n", i, a)
	}

}

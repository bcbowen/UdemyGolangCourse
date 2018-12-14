package main

import (
	"fmt"
	"strconv"
)

func main() {
	//ex01()
	ex02()
	ex03()
	//ex04()
	ex05()
	ex06()
}

func getInt() int {
	var input string
	fmt.Println("Enter a number")
	fmt.Scanln(&input)
	val, _ := strconv.Atoi(input)
	return val

}

func ex01() {
	val := getInt()
	fmt.Printf("%v %b 0x%x\n", val, val, val)
}

func ex02() {
	g := (42 == 42)
	// yadda yadda yadda... skipping rest... too simple
	fmt.Printf("g: %v", g)
}

func ex03() {

	const (
		ut        = "yadda yadda yadda"
		t  string = "I was in the pool!"
	)

	fmt.Printf("ut: %v \n", ut)
	fmt.Printf("t: %v \n", t)
}

func ex04() {
	val := getInt()
	fmt.Printf("%v %b %#x\n", val, val, val)

	val = val << 1

	fmt.Printf("%v %b %#x\n", val, val, val)
}

func ex05() {
	s := `and now for 
	somethign completely 
	different`

	fmt.Println(s)
}

func ex06() {
	const (
		next         = 2018 + iota
		postnext     = 2018 + iota
		postpostnext = 2018 + iota
		andfinally   = 2018 + iota
	)

	fmt.Println(next)
	fmt.Println(postnext)
	fmt.Println(postpostnext)
	fmt.Println(andfinally)

}

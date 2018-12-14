package main

import "fmt"

var _x int
var _y string
var _z bool

func main() {
	ex01()
	ex02()
	ex03()
	s := fmt.Sprintf("%v\t%v\t%v", _x, _y, _z)
	println(s)

	ex04()
	ex05()
}

func ex01() {
	println("ex 01")
	x := 42
	y := "JB"
	z := true

	print(x, " ", y, " ", z, "\n")
	println(x)
	println(y)
	println(z)

}

func ex02() {
	println("ex 02")

	println(_x)
	println(_y)
	println(_z)

}

func ex03() {
	println("ex 03")

	_x = 42
	_y = "yadda"
	_z = true
}

func ex04() {
	println("ex 04")

	type yadda int
	var x yadda
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}

func ex05() {
	println("ex 05")

	type yadda int
	var x yadda

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 71
	var y int = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}

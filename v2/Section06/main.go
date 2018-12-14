package main

import (
	"fmt"
	"math"
)

func main() {

	ex01()
	ex02()
	ex03()
	ex04()
	ex05()
	ex06()
	ex07()
	ex08()
	ex09()
	ex10()
}

/*

func ex() {
	fmt.Println("Ex 02 *******************************")

}

*/

func ex01() {
	fmt.Println("Ex 01 *******************************")

	f := foo()

	age, name := bar()

	fmt.Printf("%v ; %v, %v \n", f, age, name)

}

func foo() int {
	return 42

}

func bar() (int, string) {
	return 3, "yadda"

}

func ex02() {
	fmt.Println("Ex 02 *******************************")

	nums := []int{34, 23, 132, 1232, 1236, 98, 65, 54}
	foosum := foo2(nums...)
	barsum := bar2(nums)

	fmt.Printf("f: %v; b: %v \n", foosum, barsum)
}

func foo2(vals ...int) int {
	return sum(vals...)
}

func bar2(vals []int) int {
	return sum(vals...)
}

func sum(vals ...int) int {
	sum := 0
	for _, v := range vals {
		sum += v
	}

	return sum
}

func ex03() {
	fmt.Println("Ex 03 *******************************")

	f1 := func() {
		fmt.Println("f1 done")
	}

	f2 := func() {
		defer f1()
		fmt.Println("f2 through")
	}

	f2()

}

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println(p.first, p.last, p.age)

}

func ex04() {
	fmt.Println("Ex 04 *******************************")

	p := person{
		first: "Fred",
		last:  "Astaire",
		age:   34,
	}

	p.speak()

}

func ex05() {
	fmt.Println("Ex 05 *******************************")

	info := func(s shape) {
		fmt.Println("hey man, the area is", s.area())
	}

	s := square{
		side: 4,
	}

	c := circle{
		radius: 3,
	}

	info(s)
	info(c)
}

type shape interface {
	area() float64
}

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return math.Pow(s.side, 2)
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func ex06() {
	fmt.Println("Ex 06 *******************************")

	func() {
		fmt.Println("I am anonymous")
	}()
}

func ex07() {
	fmt.Println("Ex 07 *******************************")

	f := func() {
		fmt.Println("I am a function variable")
	}

	f()
}

func ex08() {
	fmt.Println("Ex 08 *******************************")

	add := func(v1 float64, v2 float64) float64 {
		return v1 + v2
	}

	getop := func() func(float64, float64) float64 {
		return add
	}

	fmt.Println("The sum is:", getop()(4.6, 34.5))
}

func ex09() {
	fmt.Println("Ex 09 *******************************")
	f := func(result int) {
		fmt.Println("And the result is: ", result)

	}

	c := func(callback func(int)) {
		callback(42)
	}

	c(f)

}

func ex10() {
	fmt.Println("Ex 10 *******************************")
	x := 0
	f := func() {
		x++
	}

	for i := 0; i < 10; i++ {
		f()
	}

	fmt.Println("Final value of x: ", x)

}

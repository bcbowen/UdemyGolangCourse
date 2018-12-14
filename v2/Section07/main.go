package main

import "fmt"

func main() {
	ex01()
	ex02()
}

func ex01() {
	x := 1
	fmt.Println(&x)
}

type person struct {
	first string
	last  string
	age   int
}

func ex02() {
	dude := person{
		first: "Fred",
		last:  "Flintstone",
		age:   34,
	}

	fmt.Println(dude)

	updatePerson(&dude, "Felix", "McGillicutty", 23)

	fmt.Println(dude)

}

func updatePerson(dude *person, first string, last string, age int) {
	(dude).first = first
	(dude).last = last
	(dude).age = age
}

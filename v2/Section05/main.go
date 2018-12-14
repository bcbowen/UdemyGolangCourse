package main

import "fmt"

func main() {
	ex01()
	ex02()
	ex03()
	ex04()
}

type person struct {
	firstName       string
	lastName        string
	favoriteFlavors []string
}

func ex01() {

	p1 := person{
		firstName:       "Jeff",
		lastName:        "Shaw",
		favoriteFlavors: []string{"hemp", "pistacio"},
	}

	p2 := person{
		firstName:       "Fred",
		lastName:        "Smith",
		favoriteFlavors: []string{"rocky road", "banana"},
	}

	fmt.Println(p1.firstName)
	fmt.Println(p1.lastName)
	for i, v := range p1.favoriteFlavors {
		fmt.Println(i, v)
	}

	fmt.Println(p2.firstName)
	fmt.Println(p2.lastName)
	for i, v := range p2.favoriteFlavors {
		fmt.Println(i, v)
	}

}

func ex02() {
	peopleMap := map[string]person{
		"Shaw": person{
			firstName:       "Jeff",
			lastName:        "Shaw",
			favoriteFlavors: []string{"hemp", "pistacio"},
		},
	}

	p1 := peopleMap["Shaw"]
	fmt.Println(p1.firstName)
	fmt.Println(p1.lastName)
	for i, v := range p1.favoriteFlavors {
		fmt.Println(i, v)
	}
}

func ex03() {
	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		vehicle
		fourWheel bool
	}

	type sedan struct {
		vehicle
		luxury bool
	}

	tacoma := truck{
		vehicle: vehicle{
			doors: 4,
			color: "silver"},

		fourWheel: true,
	}

	shitbox := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "various"},
		luxury: false,
	}

	fmt.Println(tacoma.doors)
	fmt.Println(tacoma.color)
	fmt.Println(tacoma.fourWheel)

	fmt.Println(shitbox.doors)
	fmt.Println(shitbox.color)
	fmt.Println(shitbox.luxury)
}

func ex04() {
	yadda := struct {
		name string
	}{
		name: "Costanza",
	}

	fmt.Println(yadda)
}

package main

import "fmt"

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

func ex01() {
	a := [5]int{42, 43, 44, 45, 46}

	for i, v := range a {
		fmt.Println(i, v)
	}
}

func ex02() {
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range s {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", s)
}

func ex03() {
	/*
			[42 43 44 45 46]
		[47 48 49 50 51]
		[44 45 46 47 48]
		[43 44 45 46 47]

	*/
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(s[:5])
	fmt.Println(s[5:])
	fmt.Println(s[2:7])
	fmt.Println(s[1:6])
}

func ex04() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}

func ex05() {
	/*
			start with this slice
		x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
		use APPEND & SLICING to get these values here which you should ASSIGN to a variable “y” and then print:
		[42, 43, 44, 48, 49, 50, 51]

	*/
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	var y []int

	y = append(y, x[:3]...)
	y = append(y, x[6:]...)

	fmt.Println(y)
}

func ex06() {
	states := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}

	fmt.Println(states)
	fmt.Println(len(states))
	fmt.Println(cap(states))

	for i := 0; i < len(states); i++ {
		fmt.Println(i, states[i])

	}

	/*
		for i, v := range states {
			fmt.Println(i, v)

		}
	*/
}

func ex07() {
	/*
			"James", "Bond", "Shaken, not stirred"
		"Miss", "Moneypenny", "Helloooooo, James."

	*/
	peeps := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Helloooooo, James."}}

	for i, v := range peeps {
		fmt.Println(i, v)
		for i2, v2 := range peeps[i] {
			fmt.Println(i2, v2)
		}
	}

	fmt.Println(peeps)

}

func ex08() {
	/*
			Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of TYPE []string which stores their favorite things. Store three records in your map. Print out all of the values, along with their index position in the slice.

			`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
		`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
		`no_dr`, `Being evil`, `Ice cream`, `Sunsets`

	*/
	agents := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	for k, _ := range agents {
		fmt.Println(k)
		for i, v2 := range agents[k] {
			fmt.Println("\t", i, v2)
		}
	}
}

func ex09() {
	/*
			Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of TYPE []string which stores their favorite things. Store three records in your map. Print out all of the values, along with their index position in the slice.

			`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
		`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
		`no_dr`, `Being evil`, `Ice cream`, `Sunsets`

	*/
	agents := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	agents[`bean_mr`] = []string{`puppetry`, `espionage`, `dancing`}

	for k, _ := range agents {
		fmt.Println(k)
		for i, v2 := range agents[k] {
			fmt.Println("\t", i, v2)
		}
	}
}

func ex10() {
	/*
			Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of TYPE []string which stores their favorite things. Store three records in your map. Print out all of the values, along with their index position in the slice.

			`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
		`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
		`no_dr`, `Being evil`, `Ice cream`, `Sunsets`

	*/
	agents := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	agents[`bean_mr`] = []string{`puppetry`, `espionage`, `dancing`}

	delete(agents, "no_dr")

	for k, _ := range agents {
		fmt.Println(k)
		for i, v2 := range agents[k] {
			fmt.Println("\t", i, v2)
		}
	}
}

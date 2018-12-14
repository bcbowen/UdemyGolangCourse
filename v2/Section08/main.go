package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string
	Last    string
	Sayings []string
	Age     int
}

func main() {
	ex01()
	ex02()
	ex03()
}

func ex01() {
	fmt.Println("Ex01 ************************************")
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	bs, err := json.Marshal(users)

	if err != nil {
		fmt.Println("Holy fucking shit! An error: ", err)
	}

	fmt.Println(string(bs))
}

func ex02() {
	fmt.Println("Ex02 ************************************")
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)

	var users []user
	err := json.Unmarshal([]byte(s), &users)
	if err != nil {
		fmt.Println("Holy fucking shit! An error: ", err)
	}

	fmt.Println(users)
}

func ex03() {
	fmt.Println("Ex03 ************************************")

	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	encoder := json.NewEncoder(os.Stdout)
	err := encoder.Encode(users)

	if err != nil {
		fmt.Println("Holy fucking shit! An error: ", err)
	}
}

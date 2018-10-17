package main

import (
	"flag"
	"fmt"
)

type Person struct {
	First, Last string
}

func main() {
	first := flag.String("first", "", "First Name")
	last := flag.String("last", "", "Last Name")
	flag.Parse()

	people := make(map[Person]bool)

	people[Person{"Chien", "Nguyen"}] = true
	people[Person{"Anh", "Nguyen"}] = false

	t, ok := people[Person{*first, *last}]
	if ok {
		fmt.Println(t)
		return
	}
	fmt.Println("Not present")
}

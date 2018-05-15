package main

import (
	"errors"
	"fmt"
)

func do() {
	n := 1
	defer fmt.Println("defer 1", n)

	n++
	defer fmt.Println("defer 2", n)

	n++
	defer fmt.Println("defer 3", n)
}

func checkErr() {
	var err error
	defer func() {
		if err != nil {
			fmt.Println("An error occured: ", err)
		}
	}()

	err = errors.New("Bad thing happened")
}

func main() {
	do()
	checkErr()
}

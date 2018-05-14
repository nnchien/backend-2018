package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
)

func main() {

	//read timout setup, default is zero
	delay := flag.Int("timeout", 0, "set timeout for user input")
	flag.Parse()

	customRegexp := GettingUserInput{
		label: "Input Regular Expression format: ",
	}.Start()

	regCustom, errCompile := regexp.Compile(customRegexp)
	if errCompile != nil {
		fmt.Println("Cannot compile format:", customRegexp)
		os.Exit(1)
	}

	inputTesting := GettingUserInput{
		label:     "Input for testing",
		timeInSec: *delay,
		onTimeOut: func() {
			fmt.Println("\nTimeout! You should input faster!")
			os.Exit(1)
		},
	}.Start()

	if matched := regCustom.MatchString(inputTesting); !matched {
		fmt.Printf("%s doesn't express as %s\n", inputTesting, customRegexp)
		os.Exit(1)
	}

	fmt.Println("Matched!")
}

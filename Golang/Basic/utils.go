package main

import (
	"fmt"
	"strings"
	"time"
)

//NOTE: need to call utils.go along with your target file for running

type GettingUserInput struct {
	label          string
	timeInSec      int // 0 means waiting forever
	trimmableSpace bool
	onTimeOut      TimeOutHandler
}

type TimeOutHandler func()

func (gui GettingUserInput) Start() string {
	isTimeoutCounted := gui.timeInSec > 0
	if isTimeoutCounted {
		go startTimer(gui.label, gui.timeInSec, gui.onTimeOut)
	}
	return getInput(gui.label, gui.trimmableSpace, !isTimeoutCounted)
}

func getInput(label string, trimmable bool, printable bool) string {
	if printable {
		fmt.Printf(label)
	}
	var input string
	fmt.Scanf("%s\n", &input)
	if trimmable {
		input = strings.TrimSpace(input)
	}
	return input
}

func startTimer(label string, delay int, onTimeout TimeOutHandler) {
	stop := time.After(time.Duration(delay+1) * time.Second)
	ticker := time.Tick(time.Second)
	counter := delay + 1
	for {
		select {
		case <-stop:
			onTimeout()
		case <-ticker:
			counter--
			fmt.Printf("\r%s %d(sec): ", label, counter)
		}
	}
}

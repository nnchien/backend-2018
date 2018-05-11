package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Buffer struct {
	name   string
	number int
}

func (b Buffer) process() {
	//fake a heavy work
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Proceeded", b)
}

func (b Buffer) get() {
	fmt.Println("Pushed", b)
}

var available = make(chan Buffer, 10)
var toController = make(chan Buffer)

func worker() {
	for {
		var b Buffer
		select {
		case b = <-available:
		}
		b.get()
		toController <- b
	}
}

func controller() {
	for {
		b := <-toController
		b.process()
		fmt.Print("Input: ")
	}
}

func main() {
	go worker()
	go controller()

	fmt.Println("started the machine! Press push to push and exit to stop.")
	fmt.Print("Input: ")
	reader := bufio.NewReader(os.Stdin)
	counter := 0
	text := ""

START:
	for {
		if strings.Compare(text, "exit") == 0 {
			fmt.Println("thanks for using, bye bye!")
			os.Exit(0)
		}
		text, _ = reader.ReadString('\n')
		//remove \n
		text = strings.TrimRight(text, "\n")
		if strings.Compare(text, "push") == 0 {
			available <- Buffer{name: fmt.Sprintf("event %d", counter), number: counter}
			counter++
		}
		continue START
	}
}

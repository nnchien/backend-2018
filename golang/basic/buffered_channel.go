package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Buffer struct {
	name   string
	number int
}

func (b Buffer) process() {
	//fake a heavy work
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
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
		if b.number > 0 {
			b.get()
			toController <- b
			return
		}
		fmt.Print("Push: ")
	}
}

func controller() {
	for {
		b := <-toController
		b.process()
	}
}

func main() {
	go worker()
	go controller()

	fmt.Println("started the machine! Input n (>0) to push and 0 to exit.")
	fmt.Print("Push: ")
	reader := bufio.NewReader(os.Stdin)
	text := ""

START:
	for {
		//detect if user want to exit
		if strings.Compare(text, "0") == 0 {
			fmt.Println("thanks for using, bye bye!")
			os.Exit(0)
		}

		//read input from user
		text, _ = reader.ReadString('\n')
		text = strings.TrimRight(text, "\n")

		n, err := strconv.Atoi(text)
		if err != nil {
			fmt.Print("Push: ")
			continue START
		}

		for i := 0; i < n; i++ {
			go func() {
				fmt.Println("added ", i)
				available <- Buffer{name: fmt.Sprintf("event %d", i), number: i}
			}()
		}
		continue START
	}
}

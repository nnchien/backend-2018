package main

import (
	"fmt"
	"strings"
	"sync"
)

type MyString struct {
	sync.Mutex
	str string
}

func NewMyString(s string) MyString {
	return MyString{str: s}
}

func (m MyString) OutPut() {
	m.Lock()
	defer m.Unlock()
	fmt.Println(m.str)
}

type Shouting struct {
	MyString
}

func (s Shouting) Output() {
	s.Lock()
	defer s.Unlock()
	fmt.Printf("Really loud: %s\n", s.str)
}

func NewShouting(s string) Shouting {
	load := Shouting{}
	load.str = strings.ToUpper(s)
	return load
}

func main() {
	hello := NewShouting("Hello, World!!!")
	hello.Output()
}

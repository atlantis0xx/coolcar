package main

import "fmt"

var a, b int

func f() {
	b = 2
	a = 1
}

func g() {
	fmt.Println(b) //0
	fmt.Println(a) //1
}

func main() {
	go f()
	g()
}

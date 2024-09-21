package main

import "fmt"

func main() {
	Add := Add(20, 30)
	fmt.Println(Add)
}

func Add(a, b int) int {
	return a + b
}
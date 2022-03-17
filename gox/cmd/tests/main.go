package main

import "fmt"

func PrintAny[T any](t T) {
	fmt.Println(t)
}

func main() {
	PrintAny("Hello")
	PrintAny(1)
}

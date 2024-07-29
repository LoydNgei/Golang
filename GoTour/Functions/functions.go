package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func substract(x int, y int) int {
	return x - y
}

func main() {
	fmt.Println(add(20, 45))
	fmt.Println(substract(40, 17))

}


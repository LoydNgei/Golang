package main


import "fmt"

func swap(x int, y string) (string, int) {
	return y, x
}

func main() {
	a, b := swap(5, "Hello")
	fmt.Println(a, b)
}
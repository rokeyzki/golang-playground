/*
func
method
interface
*/
package main

import "fmt"

// func
func func1() {
	fmt.Println("func1")
}

func func2(x, y int)  {
	fmt.Println("func2", x, y)
}

func func3(x, y int) int {
	return x + y
}

func func4(x, y int) (add, minus int) {
	add = x + y
	minus = x - y
	return add, minus
}

func main()  {
	// func
	fmt.Println("--------------", "func")
	func1()
	func2(1, 2)
	fmt.Println(func3(1, 2))
	fmt.Println(func4(1, 2))
}
package main

import "fmt"

var a = "菜鸟教程"
var c bool

func main() {
	fmt.Println("Hello, go!")

	// 下面这种不带声明格式的只能在函数体中出现
	b := "runoob.com"
	d, e, f := 1, 2, 3

	println(a, b, c, d, e, f)

	arr1 := [3] int {10, 20, 30}
	arr2 := [...] int {10, 20, 30, 40}
	arr3 := [2][3] int {
		{11, 12, 13},
		{21, 22, 23},
	}

	fmt.Println(arr1, arr2, arr3[0][0])
}

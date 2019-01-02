package main

import "fmt"
// interface
type Adder interface {
	add() int
}

type MyStruct struct {
	X, Y int
}

func (a *MyStruct) add() int { // method method 就是函数，只不过拥有 receiver 参数，receiver 写在 func 和 method 名字之间
	return a.X + a.Y
}

func main() {
	// pointer
	fmt.Println("--------------", "pointer")
	foo1 := 100
	var foo1p *int // pointer 声明 使用*
	foo1p = &foo1 // pointer 赋值 使用&
	fmt.Println(foo1p, *foo1p) // 使用*可以读取指针指向的值

	// array
	fmt.Println("--------------", "array")
	var arr1 [3]int // array 声明
	arr1[0] = 1 // array 赋值
	arr1[1] = 2
	fmt.Println(arr1)

	var arr2 = [3]int{1, 2, 3} // array 声明 + 赋值
	fmt.Println(arr2)

	arr3 := [...]int{10, 20, 30, 40}
	fmt.Println(arr3)

	arr4 := [2][3] int { // 多维 array
		{11, 12, 13},
		{21, 22, 23},
	}
	fmt.Println(arr4[0][0])

	// slice
	fmt.Println("--------------", "slice")
	var slice1 []int // slice 声明，空slice，值为 nil
	fmt.Println(slice1)

	var slice2 = arr2[1:2] // slice 声明 + 赋值，arr[startIndex:endIndex]
	fmt.Println(slice2)

	var slice3 = make([]int, 3, 5) // slice 声明 + 赋值，长度为3，容量为5（未必填）
	slice3[0] = 1
	slice3[1] = 2
	slice3[2] = 3
	fmt.Println("slice3 长度", len(slice3)) // len 方法获取 slice 长度
	fmt.Println("slice3 容量", cap(slice3)) // cap 方法获取 slice 容量
	slice4 := slice3[1:]
	fmt.Println(slice4)
	slice5 := append(slice4, 4)
	fmt.Println(slice5, slice4)
	slice6 := make([]int, len(slice5))
	copy(slice6, slice5)
	fmt.Println(slice6)

	slice7 := []int{1, 2, 3} // 理解容量：cap 3
	slice7 = append(slice7, 4) // cap 3*2 = 6
	slice7 = append(slice7, 5) // cap 3*2 = 6
	slice7 = append(slice7, 6) // cap 3*2 = 6
	slice7 = append(slice7, 7) // cap 6*2 = 12
	fmt.Println(slice7, len(slice7), cap(slice7))
	arrFoo := []int{1, 2, 3, 4, 5, 6}
	slice8 := arrFoo[1:4] // len:3 cap:5
	fmt.Println(slice8, len(slice8), cap(slice8))

	// map
	fmt.Println("--------------", "map")
	var map1 map[string]string
	map1 = make(map[string]string)
	map1["France"] = "巴黎"
	map1["Italy"] = "罗马"
	map1["Japan"] = "东京"
	map1["India"] = "新德里"

	fmt.Println(map1["Japan"])

	capital, ok := map1["America"]
	if (ok) {
		fmt.Println("美国的首都是", capital)
	} else {
		fmt.Println("集合中没有匹配的值")
	}
	delete(map1, "India")
	fmt.Println(map1)

	// range
	fmt.Println("--------------", "range")
	sliceRange := []int{1, 2, 3}
	sliceRangeSum := 0
	for _, num := range sliceRange {
		sliceRangeSum += num
	}
	fmt.Println("sliceRangeSum:", sliceRangeSum)

	mapRange := map[string]string{"A": "apple", "B": "banana"}
	var mapRangeStr string
	for k, v := range mapRange {
		mapRangeStr += v
		if k == "A" { mapRangeStr += " | " }
	}
	fmt.Println("mapRangeStr:", mapRangeStr)

	// struct
	fmt.Println("--------------", "struct")
	type Books struct {
		title string
		author string
		subject string
		id int
	}

	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", id: 6495407})
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})

	var book1 Books;
	book1.title = "book1 标题"
	book1.author = "book1 作者"
	book1.subject = "book1 书名"
	book1.id = 1
	fmt.Println(book1.title)

	book2 := Books{
		title: "book2 标题",
		author: "book2 作者",
		subject: "book2 书名",
		id: 2,
	}
	fmt.Println(book2.title)

	// interface
	fmt.Println("--------------", "interface")
	var f Adder
	s := MyStruct{3, 4}
	f = &s
	fmt.Println(f.add())
}

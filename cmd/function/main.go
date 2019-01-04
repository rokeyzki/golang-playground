/*
func
method
interface
package
*/
package main

import (
	"fmt"
	"time"
	pkg1 "./foo"
)

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
	return // 参数有命名的时候，return 通常可以不跟参数
}

func func5(x, y int) (minus int, error error) {
	minus = x - y
	if minus >= 0 {
		return minus, nil
	} else {
		return 0, fmt.Errorf("Error: %s", "x 必须大于或等于 y")
	}
}

func func6() func()int { // 函数 func6 返回另一个类型为 func()int 的函数
	var x int
	return func() int { // 匿名函数
		x++
		return x * x // 平方
	}
}

func func7(values... int) int { // 可变函数
	sum := 0
	for _, val := range values {
		sum += val
	}
	return sum
}

func func8() (string, int) { // 延迟函数
	foo := 100
	fmt.Println("func8 start", foo)

	defer func() { // defer 一般用于异常处理、释放资源、清理数据、记录日志等，defer 还有一个重要的特性，就是即便函数抛出了异常，也会被执行的
		time.Sleep(3 * time.Second)
		fmt.Println("func8 defer 1", foo)
	}()

	defer func() { // 按照 FILO（先进后出）的原则依次执行每一个 defer
		foo = 200
		fmt.Println("func8 defer 2", foo)
	}()

	fmt.Println("func8 middle", foo)

	return "func8 return", foo // defer 在声明时不会立即执行，而是在函数 return 执行前
}

func func9() {
	defer func() {
		if p := recover(); p != nil { // recover 捕获异常
			err := fmt.Errorf("func9 error: %v", p)
			fmt.Println(err);
		}
	}()
	panic(fmt.Sprintf("异常抛出 %q", "func9")) // panic 抛出异常
}

// method
type St1 struct {
	x, y int
}

func (s St1) Method1(z int) int { // Go 语言只有一种控制可见性的手段：大写首字母的标识符会从定义它们的包中被导出，小写字母的则不会。
	sum := s.x + s.y + z
	return sum // method 必须有 return
}

type St2 struct {
	x, y int
}

func (s *St2) method2(z int) int { // 一个struct类型的字段对同一个包的所有代码都有可见性，无论你的代码是写在一个函数还是一个方法里
	sum := s.x + s.y + z
	return sum
}

func main()  {
	// func
	fmt.Println("--------------", "func")
	func1()
	func2(1, 2)
	fmt.Println("func3", func3(1, 2))
	fmt.Println(func4(1, 2))

	func5Val, func5Err := func5(1, 2)
	if func5Err != nil {
		fmt.Println("func5", func5Err)
	} else {
		fmt.Println("func5", func5Val)
	}

	f6 := func6()
	fmt.Println("func6", f6())
	fmt.Println("func6", f6())
	fmt.Println("func6", f6())

	fmt.Println("func7", func7())
	fmt.Println("func7", func7(3))
	fmt.Println("func7", func7(1, 2, 3))

	fmt.Println(func8())

	// func9() // 要测试异常的抛出和捕获的时候再开启

	// method
	fmt.Println("--------------", "method")
	s1 := St1{1, 2}
	fmt.Println("method1", s1.Method1(3))

	s2 := &St2{1, 2}
	fmt.Println("method2", s2.method2(3))

	fmt.Println("method3", OtherMethod1()) // go run .

	// package
	fmt.Println("--------------", "package")
	fmt.Println("pkg1", pkg1.ReturnStr())
}
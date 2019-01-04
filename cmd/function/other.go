package main  // 同一个 package 目录下的方法需要一起编译才能执行：go run .

func OtherMethod1() string {
	return "from other file!"
}

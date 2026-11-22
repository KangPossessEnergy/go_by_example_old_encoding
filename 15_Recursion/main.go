package main

import "fmt"

// 1 fact函数在到达fact(0)前一直调用自身
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

// 递归
func main() {

	fmt.Println(fact(7))

	//闭包也可以是递归的，但这要求在定义闭包之前用类型化的var 显式声明闭包
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	//由于fib之前在main中声明过,因此Go知道在这里用fib调用哪个函数
	fmt.Println(fib(7))
}
